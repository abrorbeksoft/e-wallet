package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handlerv1) Hello (c *gin.Context)  {
	name:=c.Query("name")
	res:=h.Storage.Hello(name)
	c.JSON(http.StatusOK, gin.H{ "message" : res })
}
