package tag

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(router *gin.RouterGroup, database *sql.DB) {

	repository := NewRepository(database)
	service := NewService(repository)
	handler := NewHandler(service)

	financial := router.Group("/payment")

	financial.GET("/summary", handler.GetPaymentSummary)
}
