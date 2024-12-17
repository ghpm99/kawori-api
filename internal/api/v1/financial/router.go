package financial

import (
	"database/sql"
	"kawori/api/internal/api/v1/financial/payment"

	"github.com/gin-gonic/gin"
)

func RegisterFinancialRoutes(router *gin.RouterGroup, database *sql.DB) {

	financial := router.Group("/financial")
	payment.RegisterPaymentRoutes(financial, database)

}
