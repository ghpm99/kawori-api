package financial

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FinancialHandler struct {
	FinancialService *FinancialService
}

func NewFinancialHandler() *FinancialHandler {
	return &FinancialHandler{}
}

func (financialHandler *FinancialHandler) GetAllPayments(context *gin.Context) {
	payments, err := financialHandler.FinancialService.GetAllPayments()
	context.JSON(http.StatusOK, gin.H{"version": "v1", "payments": []string{payments}})
}
