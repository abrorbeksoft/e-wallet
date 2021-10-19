package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handlerv1) AllWallets(c *gin.Context)  {
	id:=c.PostForm("userId")
	res:=h.Storage.AllWallets(id)
	c.JSON(http.StatusOK, res)
}