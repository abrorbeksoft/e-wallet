package api

import (
	v1 "github.com/abrorbeksoft/e-wallet.git/api/handlers/v1"
	"github.com/abrorbeksoft/e-wallet.git/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouterOptions struct {
	Storage storage.StorageI
}

func New(options *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "e-wallet is working"})
	})


	handlerV1 :=v1.New(&v1.HandlerV1Options{
		Storage: options.Storage,
	})

	apiV1 := router.Group("/v1").Use()

	apiV1.Use()
	{
		//auth
		apiV1.POST("/login",handlerV1.Login)
		apiV1.POST("/register",handlerV1.Register)

		apiV1.GET("/hello", handlerV1.Hello)
		apiV1.POST("/allwallets",handlerV1.AllWallets)
		apiV1.POST("/getwallet",handlerV1.GetWallet)
		apiV1.POST("/addmoney",handlerV1.AddMoney)
		apiV1.POST("/removemoney",handlerV1.RemoveMoney)
		apiV1.POST("/paymenthistory",handlerV1.GetMonthlyPayment)
	}

	return router
}