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

func (h *handlerv1) GetWallet(c *gin.Context)  {
	walletId:=c.PostForm("walletId")
	res:=h.Storage.GetWallet(walletId)
	if res == nil{
		c.JSON(http.StatusOK, gin.H{"status":"error", "message" : "wallet not found"})
	} else {
		c.JSON(http.StatusOK, res)
	}
}