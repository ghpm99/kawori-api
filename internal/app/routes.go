package app

import (
	"kawori/api/internal/api/v1/financial"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	financialHandler := financial.NewFinancialHandler()
	financialGroup := router.Group("/financial")
	{
		financialGroup.GET("/", financialHandler.GetAllPayments())
	}
}
