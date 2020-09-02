package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	global "github.com/kubeinn/schutterij/internal/global"
)

func ValidatePilgrimCredentials(username string, password string) (string, error) {
	// Get pilgrim password

	// Create a JWT
	claims := &jwt.StandardClaims{
		Subject: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(global.JWT_SIGNING_KEY)
	return ss, err
}

func ValidateInnkeeperCredentials(username string, password string) (string, error) {

}

func RegisterPilgrim(username string, password string) (bool, error) {
	// Add user to database
	return true, nil
}
