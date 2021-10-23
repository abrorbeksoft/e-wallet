package v1

import (
	"github.com/abrorbeksoft/e-wallet.git/api/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

func (h *handlerv1) AddMoney(c *gin.Context)  {
	c.Param("")
	//c.ShouldBindJSON("")
	walletId := c.PostForm("walletId")
	amount := cast.ToInt64(c.PostForm("amount"))
	status,message:=h.Storage.AddMoney(walletId,amount)

	if status {
		c.JSON(http.StatusOK, gin.H{ "status" : "ok", "message" : message })
	} else {
		c.JSON(http.StatusOK, gin.H{ "status" : "error", "message" : message })
	}

}

func (h *handlerv1) RemoveMoney(c *gin.Context)  {
	walletId := c.PostForm("walletId")
	amount := cast.ToInt64(c.PostForm("amount"))
	status,message:=h.Storage.RemoveMoney(walletId,amount)
	if status {
		c.JSON(http.StatusOK, gin.H{ "status" : "ok", "message" : message })
	} else {
		c.JSON(http.StatusOK, gin.H{ "status" : "error", "message" : message })
	}
}

func (h *handlerv1) GetMonthlyPayment(c *gin.Context)  {
	walletId := c.PostForm("walletId")
	var res *models.Payments
	res=h.Storage.PaymentHistory(walletId)
	if res.Count>0 {
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, gin.H{ "status": "error", "message": "No payments in month" })
	}
}

