package routes

import (
	"kawori/api/src/routes/v1/financial"

	"github.com/gin-gonic/gin"
)

func ConfigRouter(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.GET("/financial", financial.FinancialEndpoint)
	}
}
