package app

import (
	"database/sql"
	"kawori/api/internal/api/v1/financial"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, database *sql.DB) {
	routesV1 := router.Group("/v1")
	financial.RegisterFinancialRoutes(routesV1, database)
}
