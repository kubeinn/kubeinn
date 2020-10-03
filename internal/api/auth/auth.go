package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	gin "github.com/gin-gonic/gin"
	global "github.com/kubeinn/schutterij/internal/global"
	bcrypt "golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// URL Handlers

// PostValidateCredentialsHandler is ...
func PostValidateCredentialsHandler(c *gin.Context) {
	subject := c.GetHeader("Subject")
	username := c.Query("username")
	password := c.Query("password")

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
		c.JSON(http.StatusOK, gin.H{"Authorization": jwt})

	} else if subject == "Pilgrim" {
		jwt, err := validatePilgrimCredentials(username, password)
		if err != nil {
			// Authentication failed
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid credentials provided."})
			return
		}
		// Authentication successful
		c.JSON(http.StatusOK, gin.H{"Authorization": jwt})
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid credentials provided."})
		return
	}
}

// PostRegisterPilgrim is ...
func PostRegisterPilgrim(c *gin.Context) {
	subject := c.GetHeader("Subject")
	username := c.Query("username")
	email := c.Query("email")
	password := c.Query("password")

	log.Println("===============================")
	log.Println("subject: " + subject)
	log.Println("username: " + username)
	log.Println("email: " + email)
	log.Println("password: " + password)
	log.Println("===============================")

	err := registerPilgrim(username, email, password)
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
	dbPassword, err := global.PG_CONTROLLER.SelectPilgrimPasswordByUsername(username)
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
	claims := &jwt.StandardClaims{
		Subject:  username,
		Audience: global.JWT_AUDIENCE_PILGRIM,
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
	dbPassword, err := global.PG_CONTROLLER.SelectInnkeeperPasswordByUsername(username)
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
	claims := &jwt.StandardClaims{
		Subject:  username,
		Audience: global.JWT_AUDIENCE_INNKEEPER,
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

func registerPilgrim(username string, email string, password string) error {
	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password: " + err.Error())
		return err
	}
	err = global.PG_CONTROLLER.InsertPilgrim(username, email, string(passwordHash))
	if err != nil {
		log.Println(err)
		return err
	}

	// Add user to database
	return nil
}
