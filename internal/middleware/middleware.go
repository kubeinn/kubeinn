package middleware

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	gin "github.com/gin-gonic/gin"
	global "github.com/kubeinn/schutterij/internal/global"
	"net/http"
	"strings"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// TokenAuthMiddleware is ...
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve token from header
		tokenString := c.Request.Header.Get("Authorization")
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
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		// Parse claims
		claims, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {
			// Check if sub is present in claims
			subject, ok := claims["sub"]
			if !ok {
				c.AbortWithError(http.StatusUnauthorized, errors.New("jwt does not contain the subject field"))
			}

			// Validate innkeeper privileges
			if strings.HasPrefix(urlPath, global.INNKEEPER_API_ENDPOINT_PREFIX) {
				if subject != global.JWT_SUBJECT_INNKEEPER {
					c.AbortWithError(http.StatusUnauthorized, errors.New("JWT does not contain the necessary privileges"))
				}
			}

			// Validate pilgrim privileges
			if strings.HasPrefix(urlPath, global.PILGRIM_API_ENDPOINT_PREFIX) {
				if subject == global.JWT_SUBJECT_INNKEEPER || subject == global.JWT_SUBJECT_PILGRIM {
				} else {
					c.AbortWithError(http.StatusUnauthorized, errors.New("JWT does not contain the necessary privileges"))
				}
			}
		} else {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid subject provided in JWT"))
		}
	}
}
