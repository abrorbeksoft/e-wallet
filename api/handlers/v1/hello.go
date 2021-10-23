package v1

import (
	"github.com/abrorbeksoft/e-wallet.git/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func (h *handlerv1) Hello (c *gin.Context)  {
	var res models.LoginUser;
	err:=c.ShouldBind(&res)

	if err != nil {
		h.SendBadRequest(c, err.Error())
		return
	}

	name:=c.Query("name")
	res1:=h.Storage.Hello(name)
	c.JSON(http.StatusOK, gin.H{ "message" : res1 })
}

func (h *handlerv1) SendBadRequest(c *gin.Context, message string)  {
	c.JSON(http.StatusBadRequest, gin.H{ message: message })
}