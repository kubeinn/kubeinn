package middleware

import (
	// "errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	gin "github.com/gin-gonic/gin"
	global "github.com/kubeinn/src/backend/internal/global"
	"log"
	"net/http"
	"strings"
)

// TokenAuthMiddleware is ...
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		claims := jwt.MapClaims{}
		urlPath := c.Request.URL.Path

		// Parse takes the token string and a function for looking up the key
		token, err := jwt.ParseWithClaims(tokenString, claims,
			func(token *jwt.Token) (interface{}, error) {
				// Validate the signing algorithm
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return global.JWT_SIGNING_KEY, nil
			})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
			return
		}

		// Parse claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			// Check if sub is present in claims
			audience, ok := claims["aud"]
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "JWT does not contain the audience field."})
				return
			}
			str := fmt.Sprintf("%v", claims["aud"])
			log.Println("audience: " + str)

			// Validate innkeeper privileges
			if strings.HasPrefix(urlPath, global.INNKEEPER_ROUTE_PREFIX) {
				if audience != global.JWT_AUDIENCE_INNKEEPER {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "JWT does not contain the necessary privileges."})
					return
				}
			}

			// Validate pilgrim privileges
			if strings.HasPrefix(urlPath, global.PILGRIM_ROUTE_PREFIX) {
				// Innkeepers can send requests to both innkeeper and pilgrim endpoints
				if audience != global.JWT_AUDIENCE_PILGRIM {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "JWT does not contain the necessary privileges."})
					return
				}
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Invalid audience provided in the JWT."})
			return
		}
	}
}
