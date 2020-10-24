package auth

import (
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	gin "github.com/gin-gonic/gin"
	global "github.com/kubeinn/schutterij/internal/global"
	go_cache "github.com/patrickmn/go-cache"
	bcrypt "golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// LoginRequest is...
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CustomClaims is...
type CustomClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

// URL Handlers

//PostCheckAuthHandler is ...
func PostCheckAuthHandler(c *gin.Context) {
	// Retrieve token from header
	reqToken := c.Request.Header.Get("Authorization")
	if reqToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "No Authorization header provided."})
		return
	}
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid authorization token."})
		return
	}
	tokenString := strings.TrimSpace(strings.Split(reqToken, "Bearer")[1])
	log.Println("Fetching from cache: " + tokenString)
	_, found := global.SESSION_CACHE.Get(tokenString)
	if found {
		c.JSON(http.StatusOK, gin.H{"Message": "Cache entry valid."})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Cache entry invalid."})
	}
}

// PostValidateCredentialsHandler is ...
func PostValidateCredentialsHandler(c *gin.Context) {
	subject := c.GetHeader("Subject")
	var loginRequest LoginRequest
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unable to read request body."})
		return
	}
	log.Println("body: " + string(b))
	err = json.Unmarshal(b, &loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unable to unmarshall username and password."})
		return
	}
	username := loginRequest.Username
	password := loginRequest.Password

	log.Println("===============================")
	log.Println("subject: " + subject)
	log.Println("username: " + username)
	log.Println("password: " + password)
	log.Println("===============================")

	if subject == "Innkeeper" {
		jwt, err := validateInnkeeperCredentials(username, password)
		if err != nil {
			// Authentication failed
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid credentials provided."})
			return
		}
		// Authentication successful
		log.Println("Inserting into cache: " + jwt)
		global.SESSION_CACHE.Set(jwt, "true", go_cache.DefaultExpiration)
		c.JSON(http.StatusOK, gin.H{"Authorization": jwt})

	} else if subject == "Pilgrim" {
		jwt, err := validatePilgrimCredentials(username, password)
		if err != nil {
			// Authentication failed
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid credentials provided."})
			return
		}
		// Authentication successful
		log.Println("Inserting into cache: " + jwt)
		global.SESSION_CACHE.Set(jwt, "true", go_cache.DefaultExpiration)
		c.JSON(http.StatusOK, gin.H{"Authorization": jwt})
	} else if subject == "Reeve" {
		jwt, err := validateReeveCredentials(username, password)
		if err != nil {
			// Authentication failed
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid credentials provided."})
			return
		}
		// Authentication successful
		log.Println("Inserting into cache: " + jwt)
		global.SESSION_CACHE.Set(jwt, "true", go_cache.DefaultExpiration)
		c.JSON(http.StatusOK, gin.H{"Authorization": jwt})
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid credentials provided."})
		return
	}
}

// PostRegisterPilgrim is ...
func PostRegisterPilgrim(c *gin.Context) {
	subject := c.GetHeader("Subject")
	vic := c.Query("vic")
	username := c.Query("username")
	email := c.Query("email")
	password := c.Query("password")

	log.Println("===============================")
	log.Println("subject: " + subject)
	log.Println("vic: " + vic)
	log.Println("username: " + username)
	log.Println("email: " + email)
	log.Println("password: " + password)
	log.Println("===============================")

	err := RegisterPilgrim(vic, username, email, password)
	if err != nil {
		// Registration failed
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Registration error."})
		return
	}
	// Registration successful
	c.JSON(http.StatusOK, gin.H{"Message": "User registered successfully."})
}

// PostRegisterVillage is ...
func PostRegisterVillage(c *gin.Context) {
	subject := c.GetHeader("Subject")
	organization := c.Query("organization")
	description := c.Query("description")
	username := c.Query("username")
	email := c.Query("email")
	password := c.Query("password")

	log.Println("===============================")
	log.Println("subject: " + subject)
	log.Println("organization: " + organization)
	log.Println("description: " + description)
	log.Println("username: " + username)
	log.Println("email: " + email)
	log.Println("password: " + password)
	log.Println("===============================")

	err := RegisterVillage(organization, description)
	if err != nil {
		// Registration failed
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Registration error."})
		return
	}

	err = RegisterReeve(organization, username, email, password)
	if err != nil {
		// Registration failed
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Registration error."})
		return
	}
	// Registration successful
	c.JSON(http.StatusOK, gin.H{"Message": "User registered successfully."})
}

// Internal functions
func validatePilgrimCredentials(username string, password string) (string, error) {
	// Get password from database
	dbID, dbPassword, err := global.PG_CONTROLLER.SelectPilgrimByUsername(username)
	if err != nil {
		log.Println(err)
	}
	log.Println("dbID: " + string(dbID))
	log.Println("dbPassword: " + dbPassword)
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		log.Println("Failed to authenticate user: " + username)
		return "", err
	}
	log.Println("Successfully authenticated user: " + username)

	// Password matches, proceed to create a JWT
	claims := CustomClaims{
		dbID,
		jwt.StandardClaims{
			Subject:   dbID,
			Audience:  global.JWT_AUDIENCE_PILGRIM,
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(global.JWT_SIGNING_KEY)
	if err != nil {
		log.Println("Failed to generate JWT: " + err.Error())
		return "", err
	}
	log.Println("JWT: " + ss)
	return ss, nil
}

func validateReeveCredentials(username string, password string) (string, error) {
	// Get password from database
	dbID, dbPassword, err := global.PG_CONTROLLER.SelectReeveByUsername(username)
	if err != nil {
		log.Println(err)
	}
	log.Println("dbID: " + string(dbID))
	log.Println("dbPassword: " + dbPassword)
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		log.Println("Failed to authenticate user: " + username)
		return "", err
	}
	log.Println("Successfully authenticated user: " + username)

	// Password matches, proceed to create a JWT
	claims := CustomClaims{
		dbID,
		jwt.StandardClaims{
			Subject:   dbID,
			Audience:  global.JWT_AUDIENCE_REEVE,
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(global.JWT_SIGNING_KEY)
	if err != nil {
		log.Println("Failed to generate JWT: " + err.Error())
		return "", err
	}
	log.Println("JWT: " + ss)
	return ss, nil
}

func validateInnkeeperCredentials(username string, password string) (string, error) {
	// Get password from database
	dbID, dbPassword, err := global.PG_CONTROLLER.SelectInnkeeperByUsername(username)
	if err != nil {
		log.Println(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		log.Println("Failed to authenticate user: " + username)
		return "", err
	}
	log.Println("Successfully authenticated user: " + username)

	// Password matches, proceed to create a JWT
	claims := CustomClaims{
		"postgres",
		jwt.StandardClaims{
			Subject:   dbID,
			Audience:  global.JWT_AUDIENCE_INNKEEPER,
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(global.JWT_SIGNING_KEY)
	if err != nil {
		log.Println("Failed to generate JWT: " + err.Error())
		return "", err
	}
	log.Println("JWT: " + ss)
	return ss, nil
}

func RegisterPilgrim(vic string, username string, email string, password string) error {
	villageID, err := global.PG_CONTROLLER.SelectVillageByVIC(vic)
	if err != nil {
		log.Println("Failed to retrieve corresponding villageID: " + err.Error())
		return err
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password: " + err.Error())
		return err
	}

	// Add user to database
	err = global.PG_CONTROLLER.InsertPilgrim(username, email, string(passwordHash), villageID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func RegisterVillage(organization string, description string) error {
	// Add village to database
	err := global.PG_CONTROLLER.InsertVillage(organization, description)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Test
// RegisterInnkeeper is ...
func RegisterInnkeeper(username string, email string, password string) error {
	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password: " + err.Error())
		return err
	}
	// Add user to database
	err = global.PG_CONTROLLER.InsertInnkeeper(username, email, string(passwordHash))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func RegisterReeve(organization string, username string, email string, password string) error {
	villageID, err := global.PG_CONTROLLER.SelectVillageByOrganization(organization)
	if err != nil {
		log.Println("Failed to retrieve corresponding villageID: " + err.Error())
		return err
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password: " + err.Error())
		return err
	}
	// Add user to database
	err = global.PG_CONTROLLER.InsertReeve(username, email, string(passwordHash), villageID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
