package payment

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(router *gin.RouterGroup, database *sql.DB) {

	repository := NewRepository(database)
	service := NewService(repository)
	handler := NewHandler(service)

	financial := router.Group("/payment")

	financial.GET("/summary/", handler.GetPaymentSummaryHandler)
	financial.GET("/", handler.GetAllPaymentHandler)
	financial.GET("/:paymentid/", handler.GetPaymentHandler)
	financial.POST("/:paymentid/save", handler.SavePaymentHandler)
	financial.POST("/:paymentid/payoff", handler.PayoffPaymentHandler)
}
