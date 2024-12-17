package contract

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(router *gin.RouterGroup, database *sql.DB) {

	repository := NewRepository(database)
	service := NewService(repository)
	handler := NewHandler(service)

	financial := router.Group("/contract")

	financial.GET("/", handler.GetPaymentSummary)
}
