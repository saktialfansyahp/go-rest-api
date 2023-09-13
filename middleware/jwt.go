package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/saktialfansyahp/go-rest-api/config"
	"github.com/saktialfansyahp/go-rest-api/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if  err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusOK, response)
				return
			}
		}

		tokenString := c.Value

		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				response := map[string]string{"message": "Unauthorized, Token expired!"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GINMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// cookie, err := c.Cookie("token")
		// if err != nil {
		// 	if err == http.ErrNoCookie {
		// 		c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		// 		c.Abort()
		// 		return
		// 	}
		// }
		authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
            c.Abort()
            return
        }

        // Split header value to extract the token
        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
            c.Abort()
            return
        }

        // Extract and validate the token
        tokenString := tokenParts[1]
		
		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
				c.Abort()
				return
			case jwt.ValidationErrorExpired:
				c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized, Token Expired!"})
				c.Abort()
				return
			default:
				c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
				c.Abort()
				return
			}
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			c.Abort()
			return
		}

		authorized := false
		for _, requiredRole := range requiredRoles {
			if requiredRole == claims.Role {
				authorized = true
				break
			}
		}

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"Message": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}