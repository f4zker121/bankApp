package handler

import (
	"bankApp/pkg/account"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	accounts *account.Account
}

func NewHandler(accounts *account.Account) *Handler {
	return &Handler{accounts: accounts}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	accounts := router.Group("/accounts")
	{
		accounts.POST("/", h.createAcc)
		accounts.POST("/:id/deposit", h.deposit)
		accounts.POST("/:id/withdraw", h.withdraw)
		accounts.GET("/:id/balance", h.getBalance)
	}

	return router
}
