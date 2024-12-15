package financial

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(financialService *Service) *Handler {
	return &Handler{service: financialService}
}

func (handler *Handler) GetPaymentSummary(context *gin.Context) {

	pageQuery := context.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	pageSizeQuery := context.DefaultQuery("page_size", "15")
	pageSize, err := strconv.Atoi(pageSizeQuery)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	dataInicialQuery := context.Query("data_inicial")
	dataInicial, err := time.Parse("2006-01-02", dataInicialQuery)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	dataFinalQuery := context.Query("data_final")

	dataFinal, err := time.Parse("2006-01-02", dataFinalQuery)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	payments, err := handler.service.GetPaymentSummary(Pagination{
		Page:     page,
		PageSize: pageSize,
	}, PaymentSummaryFilter{
		UserId:      1,
		DataInicial: dataInicial,
		DataFinal:   dataFinal,
	})
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
	}
	context.JSON(http.StatusOK, gin.H{"version": "v1", "payments": payments})
}
