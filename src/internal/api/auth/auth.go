package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	gin "github.com/gin-gonic/gin"
	global "github.com/kubeinn/kubeinn/src/internal/global"
	go_cache "github.com/patrickmn/go-cache"
	bcrypt "golang.org/x/crypto/bcrypt"
)

// LoginRequest represents the HTTP request when a user attempts to log in
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CustomClaims represents the structure of the JSON Web Token claims
type CustomClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

// PostCheckAuthHandler handles POST requests to the /check-auth authentication API endpoint.
// It checks if the request contains the Authorization token.
// If the token exists and is valid, this handler responds with a StatusOK response code.
// Otherwise, this handler responds with a StatusUnauthorized response code.
func PostCheckAuthHandler(c *gin.Context) {
	// Retrieve token from header
	reqToken := c.Request.Header.Get("Authorization")
	if reqToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No Authorization header provided."})
		return
	}

	// Retrieve Bearer token
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization token."})
		return
	}
	tokenString := strings.TrimSpace(strings.Split(reqToken, "Bearer")[1])

	// Retrieve token from cache
	_, found := global.SESSION_CACHE.Get(tokenString)
	if found {
		c.JSON(http.StatusOK, gin.H{"message": "Cache entry valid."})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Cache entry invalid."})
	}
}

// PostValidateCredentialsHandler handles POST requests to the /login authentication API endpoint.
// It compares the username and password provided to the credentials stored in the database.
// If a match is found, this handler responds with a StatusOK response code
// and an authorization token with the corresponding role.
// Otherwise, this handler responds with a StatusUnauthorized response code.
func PostValidateCredentialsHandler(c *gin.Context) {
	// Instantiate login request
	var loginRequest LoginRequest

	// Retrieve subject header
	subject := c.GetHeader("Subject")

	// Read in request body
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unable to read request body."})
		return
	}

	// Unmarshall request body into login request structure
	err = json.Unmarshal(b, &loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unable to unmarshall username and password."})
		return
	}

	// Retrieve username and password
	username := loginRequest.Username
	password := loginRequest.Password

	// Check if subject is innkeeper or pilgrim
	if subject == global.JWT_AUDIENCE_INNKEEPER {
		// Check if credentials provided matches that of a cluster administrator's

		// Get password from database
		dbID, dbPassword, err := global.PG_CONTROLLER.SelectInnkeeperByUsername(username)
		if err != nil {
			log.Println(err)
		}

		// Compare hash of password provided
		err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid username and password provided."})
			return
		}

		// Authentication successful
		log.Println("Successfully authenticated user: " + username)

		// Password matches, proceed to create a JWT
		claims := CustomClaims{
			dbID,
			jwt.StandardClaims{
				Subject:   dbID,
				Audience:  global.JWT_AUDIENCE_INNKEEPER,
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(global.JWT_SIGNING_KEY)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Failed to generate JWT."})
			return
		}
		// Authentication successful

		// Add JWT to cache
		global.SESSION_CACHE.Set(ss, "true", go_cache.DefaultExpiration)

		// Return StatusOK response code with authorization token
		c.JSON(http.StatusOK, gin.H{"Authorization": ss})
	} else if subject == global.JWT_AUDIENCE_PILGRIM {
		// Check if credentials provided matches that of a cluster user's

		// Get password from database
		dbID, dbPassword, status, err := global.PG_CONTROLLER.SelectPilgrimByUsername(username)
		if err != nil {
			// Authentication failed
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No user by that username found."})
			return
		}

		// Check if status is accepted
		if status != "accepted" {
			// Authentication failed
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Status of pilgrim account is not accepted."})
			return
		}

		// Compare hash of password provided
		err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid username and password provided."})
			return
		}

		// Authentication successful
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Failed to generate JWT."})
			return
		}

		// Add JWT to cache
		global.SESSION_CACHE.Set(ss, "true", go_cache.DefaultExpiration)

		// Return StatusOK response code with authorization token
		c.JSON(http.StatusOK, gin.H{"Authorization": ss})
	} else {
		// Subject is neither innkeeper nor pilgrim
		// Return StatusUnauthorized response code with error message
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials provided."})
		return
	}
}

// PostRegisterPilgrimHandler handles POST requests to the /register/pilgrim authentication API endpoint.
// It hashes the password provided and stores the details of the new user in the database.
// If registration is successful, this handler responds with a StatusOK response code.
// Otherwise, this handler responds with a StatusBadRequest response code.
func PostRegisterPilgrimHandler(c *gin.Context) {
	// Retrieve query perimeters
	organization := c.Query("organization")
	description := c.Query("description")
	username := c.Query("username")
	email := c.Query("email")
	password := c.Query("password")

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Registration error."})
		return
	}

	// Add cluster user to database
	err = global.PG_CONTROLLER.InsertPilgrim(organization, description, username, email, string(passwordHash))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Registration error."})
		return
	}

	// Registration successful
	c.JSON(http.StatusOK, gin.H{"message": "Pilgrim request submitted successfully."})
	return
}

// RegisterInnkeeper registers an innkeeper account into the database
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
