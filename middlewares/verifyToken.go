package middlewares

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Middleware struct{}

func (m *Middleware) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* Get Token From Header */
		json := []byte(`{"token": "` + c.Request.Header.Get("Authorization") + `"}`)
		/* Call API Verify Token */
		// response, _ := http.Post("http://farmer-backend-test:9000/api/v1/verify", "application/json", bytes.NewBuffer(json))
		response, _ := http.Post("http://localhost:9000/api/v1/verify", "application/json", bytes.NewBuffer(json))
		/* Check Response */
		if response.StatusCode == 200 {
			req_token := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", -1)
			token, _, err := jwt.NewParser().ParseUnverified(req_token, jwt.MapClaims{})
			payload := token.Claims.(jwt.MapClaims)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"success": false,
					"message": "Unauthorized",
				})
			}
			c.Request.Header.Set("user_id", payload["user_id"].(string))
			c.Request.Header.Set("username", payload["username"].(string))
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
		}
	}
}
