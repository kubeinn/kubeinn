package hooks

import (
	"encoding/json"
	"errors"
	gin "github.com/gin-gonic/gin"
	global "github.com/kubeinn/src/backend/internal/global"
	bcrypt "golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"strings"
)

type InnkeeperCreateRequestBody struct {
	Username string `json:"username"`
	Password string `json:"passwd"`
	Email    string `json:"email"`
}

type InnkeeperEditRequestBody struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"passwd"`
	Email    string `json:"email"`
}

type PilgrimCreateRequestBody struct {
	Organization string `json:"organization"`
	Description  string `json:"description"`
	Username     string `json:"username"`
	Password     string `json:"passwd"`
	Email        string `json:"email"`
	Status       string `json:"status"`
}

type PilgrimEditRequestBody struct {
	ID           string `json:"id"`
	Organization string `json:"organization"`
	Description  string `json:"description"`
	Username     string `json:"username"`
	Password     string `json:"passwd"`
	Email        string `json:"email"`
	Status       string `json:"status"`
}

type ProjectCreateRequestBody struct {
	PilgrimID string `json:"pilgrimid"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	CPU       int64  `json:"cpu"`
	Memory    int64  `json:"memory"`
	Storage   int64  `json:"storage"`
}

// PreCreateProjectHook is ...
func PreCreateProjectHook(c *gin.Context, audience string) ([]byte, error) {
	var projectCreateRequestBody ProjectCreateRequestBody

	log.Println("Decoding JSON...")
	err := json.NewDecoder(c.Request.Body).Decode(&projectCreateRequestBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	namespace := projectCreateRequestBody.Title
	cpu := projectCreateRequestBody.CPU
	memory := projectCreateRequestBody.Memory
	storage := projectCreateRequestBody.Storage
	if err != nil {
		return nil, err
	}
	err = global.KUBE_CONTROLLER.CreateNamespace(namespace)
	if err != nil {
		return nil, err
	}

	err = global.KUBE_CONTROLLER.CreateResourceQuota(namespace, cpu, memory, storage)
	if err != nil {
		return nil, err
	}

	err = global.KUBE_CONTROLLER.CreateServiceAccount(namespace)
	if err != nil {
		return nil, err
	}

	err = global.KUBE_CONTROLLER.CreateRole(namespace)
	if err != nil {
		return nil, err
	}

	err = global.KUBE_CONTROLLER.CreateRoleBinding(namespace)
	if err != nil {
		return nil, err
	}

	kubecfg, err := global.KUBE_CONTROLLER.GenerateKubeConfiguration(namespace)
	if err != nil {
		return nil, err
	}

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
	body, err := json.Marshal(newReqBody)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// PreDeleteProjectHook is ...
func PreDeleteProjectHook(c *gin.Context, audience string, subject string) error {
	id := strings.TrimPrefix(c.Query("id"), "eq.")

	// Get title from database
	dbPilgrimID, dbTitle, err := global.PG_CONTROLLER.SelectProjectById(id)
	if err != nil {
		return err
	}

	if audience != global.JWT_AUDIENCE_INNKEEPER {
		if dbPilgrimID != subject {
			log.Println("invalid subject: " + dbPilgrimID)
			return errors.New("invalid subject")
		}
	}

	err = global.KUBE_CONTROLLER.DeleteNamespace(dbTitle)
	if err != nil {
		return err
	}
	return nil
}

// PreCreateInnkeeperHook is ...
func PreCreateInnkeeperHook(c *gin.Context) ([]byte, error) {
	var innkeeperCreateRequestBody InnkeeperCreateRequestBody

	log.Println("Decoding JSON...")
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

// PreEditInnkeeperHook is ...
func PreEditInnkeeperHook(c *gin.Context) ([]byte, error) {
	var innkeeperEditRequestBody InnkeeperEditRequestBody

	log.Println("Decoding JSON...")
	err := json.NewDecoder(c.Request.Body).Decode(&innkeeperEditRequestBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	newReqBody := make(map[string]string)
	newReqBody["id"] = innkeeperEditRequestBody.ID
	newReqBody["username"] = innkeeperEditRequestBody.Username
	newReqBody["email"] = innkeeperEditRequestBody.Email
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

	body, err := json.Marshal(newReqBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}

// PreCreatePilgrimHook is ...
func PreCreatePilgrimHook(c *gin.Context) ([]byte, error) {
	var pilgrimCreateRequestBody PilgrimCreateRequestBody

	log.Println("Decoding JSON...")
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

// PreEditPilgrimHook is ...
func PreEditPilgrimHook(c *gin.Context) ([]byte, error) {
	var pilgrimEditRequestBody PilgrimEditRequestBody

	log.Println("Decoding JSON...")
	err := json.NewDecoder(c.Request.Body).Decode(&pilgrimEditRequestBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	newReqBody := make(map[string]string)
	newReqBody["id"] = pilgrimEditRequestBody.ID
	newReqBody["organization"] = pilgrimEditRequestBody.Organization
	newReqBody["description"] = pilgrimEditRequestBody.Description
	newReqBody["username"] = pilgrimEditRequestBody.Username
	newReqBody["email"] = pilgrimEditRequestBody.Email
	newReqBody["status"] = pilgrimEditRequestBody.Status
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
	body, err := json.Marshal(newReqBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}
