package contract

import (
	"kawori/api/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(financialService *Service) *Handler {
	return &Handler{service: financialService}
}

func (handler *Handler) GetPaymentSummary(context *gin.Context) {
	userContext, exist := context.Get("user")

	if !exist {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Este usuário não possui permissão para acessar este módulo.",
		})
	}

	userData := userContext.(utils.User)

	page := utils.ParseInt(context.DefaultQuery("page", "1"), context)
	pageSize := utils.ParseInt(context.DefaultQuery("page_size", "15"), context)

	startDate := utils.ParseDate(context.Query("start_date"), context)
	endDate := utils.ParseDate(context.Query("end_date"), context)

	payments, err := handler.service.GetPaymentSummary(
		Pagination{
			Page:     page,
			PageSize: pageSize,
		}, PaymentSummaryFilter{
			UserId:    userData.Id,
			StartDate: startDate,
			EndDate:   endDate,
		},
	)
	if err != nil {
		log.Println(err)
		context.AbortWithError(http.StatusInternalServerError, err)
	}

	context.JSON(http.StatusOK, gin.H{
		"payments":  payments.data,
		"page_info": payments.pageInfo,
	})
}
