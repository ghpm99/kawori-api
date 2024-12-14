package financial

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(financialService *Service) *Handler {
	return &Handler{service: financialService}
}

func (financialHandler *Handler) GetAllPayments(context *gin.Context) {
	payments, err := financialHandler.service.GetAllPayments()
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusOK, gin.H{"version": "v1", "payments": payments})
}
