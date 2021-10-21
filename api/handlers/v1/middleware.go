package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token:=c.GetHeader("X-UserId")
		if token == "" {
			c.JSON(http.StatusOK,gin.H{
				"message" : "Header not found",
			})
			c.Abort()
			return
		}
		Claims,err:=ExtractClaims(token)
		if err!=nil {
			c.JSON(http.StatusOK, gin.H{ "message": "Error while extracting the claims" })
			c.Abort()
			return
		}
		c.Set("userId",Claims["id"])
		c.Set("username", Claims["login"])
		c.Next()
	}
}
