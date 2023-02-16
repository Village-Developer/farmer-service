package middlewares

import (
	"bytes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func (m *Middleware) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* Get Token From Header */
		json := []byte(`{"token": "` + c.Request.Header.Get("Authorization") + `"}`)
		/* Call API Verify Token */
		response, err := http.Post("http://localhost:9000/api/v1/verify", "application/json", bytes.NewBuffer(json))
		/* Check Response */
		if response.StatusCode != 200 || err != nil {
			log.Println(response.StatusCode, err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
		} else {
			log.Println(response.StatusCode)
			c.Next()
		}
	}
}
