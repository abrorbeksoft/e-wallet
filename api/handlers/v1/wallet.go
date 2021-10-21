package v1

import (
	"github.com/abrorbeksoft/e-wallet.git/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handlerv1) AllWallets(c *gin.Context)  {
	id:=c.MustGet("userId").(string)
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

func (h *handlerv1) CreateWallet(c *gin.Context) {
	var (
		id string
	)
	id=c.MustGet("userId").(string)
	walletType:=c.PostForm("type")
	if walletType=="" {
		c.JSON(http.StatusOK, gin.H{
			"message" : "Please fill required fields",
		})
		return
	}

	res, err:=h.Storage.CreateWallet(&models.CreateWallet{
		Type: walletType,
	}, id)

	if err!=nil {
		c.JSON(http.StatusOK, gin.H{
			"message" : "Error occured while creating wallet",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"walletId" : res,
		"message" : "ok",
	})
}