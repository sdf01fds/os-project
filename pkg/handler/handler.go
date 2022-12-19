package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sdf0106/os-project/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", h.signUp)
		auth.POST("/sign_in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		wallet := api.Group("/wallet")
		{
			wallet.GET("/", h.getAllWallets)
			wallet.GET("/:id", h.getWalletById)
			wallet.POST("/", h.createWallet)
			wallet.PUT("/:id", h.updateWalletBalance)

		}
		transactions := api.Group("/transactions")
		{
			transactions.GET("/", h.getAllTransactions)
			transactions.GET("/:id", h.getTransactionById)
			transactions.POST("/", h.createTransaction)
			transactions.DELETE("/:id", h.deleteTransaction)
		}
	}

	return router
}
