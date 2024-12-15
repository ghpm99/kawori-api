package financial

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterFinancialRoutes(router *gin.RouterGroup, database *sql.DB) {

	repository := NewRepository(database)
	service := NewService(repository)
	handler := NewHandler(service)

	financial := router.Group("/financial")

	financial.GET("/", handler.GetPaymentSummary)
}
