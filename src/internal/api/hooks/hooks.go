package hooks

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"

	gin "github.com/gin-gonic/gin"
	global "github.com/kubeinn/kubeinn/src/internal/global"
	bcrypt "golang.org/x/crypto/bcrypt"
)

// InnkeeperCreateRequestBody represents the HTTP request when a innkeeper is to be created
type InnkeeperCreateRequestBody struct {
	Username string `json:"username"`
	Password string `json:"passwd"`
	Email    string `json:"email"`
}

// InnkeeperEditRequestBody represents the HTTP request when details of an innkeeper are to be changed
type InnkeeperEditRequestBody struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"passwd"`
	Email    string `json:"email"`
}

// PilgrimCreateRequestBody represents the HTTP request when a pilgrim is to be created
type PilgrimCreateRequestBody struct {
	Organization string `json:"organization"`
	Description  string `json:"description"`
	Username     string `json:"username"`
	Password     string `json:"passwd"`
	Email        string `json:"email"`
	Status       string `json:"status"`
}

// PilgrimEditRequestBody represents the HTTP request when details of an pilgrim are to be changed
type PilgrimEditRequestBody struct {
	ID           string `json:"id"`
	Organization string `json:"organization"`
	Description  string `json:"description"`
	Username     string `json:"username"`
	Password     string `json:"passwd"`
	Email        string `json:"email"`
	Status       string `json:"status"`
}

// ProjectCreateRequestBody represents the HTTP request when a project is to be created
type ProjectCreateRequestBody struct {
	PilgrimID string `json:"pilgrimid"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	CPU       int64  `json:"cpu"`
	Memory    int64  `json:"memory"`
	Storage   int64  `json:"storage"`
}

// PreCreateProjectHook is called before a project is created.
// This hook creates objects in the  cluster corresponding to the create project request.
// After objects are created, a new request body is returned which will be sent to the client.
func PreCreateProjectHook(c *gin.Context, audience string) ([]byte, error) {
	// Instantiate project create request body
	var projectCreateRequestBody ProjectCreateRequestBody

	// Decode JSON
	err := json.NewDecoder(c.Request.Body).Decode(&projectCreateRequestBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Retrieve parameters from projectCreateRequestBody
	namespace := projectCreateRequestBody.Title
	cpu := projectCreateRequestBody.CPU
	memory := projectCreateRequestBody.Memory
	storage := projectCreateRequestBody.Storage
	if err != nil {
		return nil, err
	}

	// Create namespace
	err = global.KUBE_CONTROLLER.CreateNamespace(namespace)
	if err != nil {
		return nil, err
	}

	// Create resource quota
	err = global.KUBE_CONTROLLER.CreateResourceQuota(namespace, cpu, memory, storage)
	if err != nil {
		return nil, err
	}

	// Create service account
	err = global.KUBE_CONTROLLER.CreateServiceAccount(namespace)
	if err != nil {
		return nil, err
	}

	// Create role
	err = global.KUBE_CONTROLLER.CreateRole(namespace)
	if err != nil {
		return nil, err
	}

	// Create role binding
	err = global.KUBE_CONTROLLER.CreateRoleBinding(namespace)
	if err != nil {
		return nil, err
	}

	// Create network policy
	err = global.KUBE_CONTROLLER.CreateNetworkPolicy(namespace)
	if err != nil {
		return nil, err
	}

	// Create kube configuration file
	kubecfg, err := global.KUBE_CONTROLLER.GenerateKubeConfiguration(namespace)
	if err != nil {
		return nil, err
	}

	// Create a request body and format with kubernetes objects created
	newReqBody := make(map[string]string)
	if audience == global.JWT_AUDIENCE_INNKEEPER {
		newReqBody["pilgrimid"] = projectCreateRequestBody.PilgrimID
	}
	newReqBody["title"] = projectCreateRequestBody.Title
	newReqBody["details"] = projectCreateRequestBody.Details
	newReqBody["cpu"] = strconv.FormatInt(projectCreateRequestBody.CPU, 10)
	newReqBody["memory"] = strconv.FormatInt(projectCreateRequestBody.Memory, 10)
	newReqBody["storage"] = strconv.FormatInt(projectCreateRequestBody.Storage, 10)
	newReqBody["kube_configuration"] = kubecfg

	// Marshal request body
	body, err := json.Marshal(newReqBody)
	if err != nil {
		return nil, err
	}

	// Return request body
	return body, nil
}

// PreDeleteProjectHook is called before project is deleted.
// This hook deletes objects in the  cluster corresponding to the delete project request.
func PreDeleteProjectHook(c *gin.Context, audience string, subject string) error {
	// Parse id
	id := strings.TrimPrefix(c.Query("id"), "eq.")

	// Get title from database
	dbPilgrimID, dbTitle, err := global.PG_CONTROLLER.SelectProjectByID(id)
	if err != nil {
		return err
	}

	// Check if user has the privileges to delete project
	if audience != global.JWT_AUDIENCE_INNKEEPER {
		if dbPilgrimID != subject {
			log.Println("invalid subject: " + dbPilgrimID)
			return errors.New("invalid subject")
		}
	}

	// Delete the project
	err = global.KUBE_CONTROLLER.DeleteNamespace(dbTitle)
	if err != nil {
		return err
	}
	return nil
}

// PreCreateInnkeeperHook is called before a innkeeper is created.
// This hook prepares the proxy request to be sent to PostgREST.
// After request is formatted, it is returned and sent to the PostgREST API.
func PreCreateInnkeeperHook(c *gin.Context) ([]byte, error) {
	// Instantiate innkeeperCreateRequestBody
	var innkeeperCreateRequestBody InnkeeperCreateRequestBody

	// Decode JSON
	err := json.NewDecoder(c.Request.Body).Decode(&innkeeperCreateRequestBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Hash password
	log.Println("Hashing password...")
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(innkeeperCreateRequestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create request body
	newReqBody := make(map[string]string)
	newReqBody["username"] = innkeeperCreateRequestBody.Username
	newReqBody["email"] = innkeeperCreateRequestBody.Email
	newReqBody["passwd"] = string(passwordHash)
	body, err := json.Marshal(newReqBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}

// PreEditInnkeeperHook is called before a innkeeper is edited.
// This hook prepares the proxy request to be sent to PostgREST.
// After request is formatted, it is returned and sent to the PostgREST API.
func PreEditInnkeeperHook(c *gin.Context) ([]byte, error) {
	// Instantiate innkeeperEditRequestBody
	var innkeeperEditRequestBody InnkeeperEditRequestBody

	// Decode JSON
	err := json.NewDecoder(c.Request.Body).Decode(&innkeeperEditRequestBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create new request body
	newReqBody := make(map[string]string)
	newReqBody["id"] = innkeeperEditRequestBody.ID
	newReqBody["username"] = innkeeperEditRequestBody.Username
	newReqBody["email"] = innkeeperEditRequestBody.Email

	// Check if there are any changes to password
	// If none, return empty string
	if innkeeperEditRequestBody.Password != "" {
		// Hash password
		log.Println("Hashing password...")
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(innkeeperEditRequestBody.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		newReqBody["passwd"] = string(passwordHash)
	} else {
		newReqBody["passwd"] = ""
	}

	// Marshal request body
	body, err := json.Marshal(newReqBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}

// PreCreatePilgrimHook is called before a pilgrim is created.
// This hook prepares the proxy request to be sent to PostgREST.
// After request is formatted, it is returned and sent to the PostgREST API.
func PreCreatePilgrimHook(c *gin.Context) ([]byte, error) {
	// Instantiate pilgrimCreateRequestBody
	var pilgrimCreateRequestBody PilgrimCreateRequestBody

	// Decode JSON
	err := json.NewDecoder(c.Request.Body).Decode(&pilgrimCreateRequestBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Hash password
	log.Println("Hashing password...")
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pilgrimCreateRequestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create new request body
	newReqBody := make(map[string]string)
	newReqBody["organization"] = pilgrimCreateRequestBody.Organization
	newReqBody["description"] = pilgrimCreateRequestBody.Description
	newReqBody["username"] = pilgrimCreateRequestBody.Username
	newReqBody["email"] = pilgrimCreateRequestBody.Email
	newReqBody["passwd"] = string(passwordHash)
	newReqBody["status"] = pilgrimCreateRequestBody.Status
	body, err := json.Marshal(newReqBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}

// PreEditPilgrimHook is called before a pilgrim is edited.
// This hook prepares the proxy request to be sent to PostgREST.
// After request is formatted, it is returned and sent to the PostgREST API.
func PreEditPilgrimHook(c *gin.Context) ([]byte, error) {
	// Instantiate pilgrimEditRequestBody
	var pilgrimEditRequestBody PilgrimEditRequestBody

	// Decode JSON
	err := json.NewDecoder(c.Request.Body).Decode(&pilgrimEditRequestBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create new request body
	newReqBody := make(map[string]string)
	newReqBody["id"] = pilgrimEditRequestBody.ID
	newReqBody["organization"] = pilgrimEditRequestBody.Organization
	newReqBody["description"] = pilgrimEditRequestBody.Description
	newReqBody["username"] = pilgrimEditRequestBody.Username
	newReqBody["email"] = pilgrimEditRequestBody.Email
	newReqBody["status"] = pilgrimEditRequestBody.Status

	// Check if there are any changes to password
	// If none, return empty string
	if pilgrimEditRequestBody.Password != "" {
		// Hash password
		log.Println("Hashing password...")
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(pilgrimEditRequestBody.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		newReqBody["passwd"] = string(passwordHash)
	} else {
		newReqBody["passwd"] = ""
	}

	// Marshal request body
	body, err := json.Marshal(newReqBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}
