package payment

import (
	"fmt"
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

func (handler *Handler) GetPaymentSummaryHandler(context *gin.Context) {
	userContext, exist := context.Get("user")

	if !exist {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Este usuário não possui permissão para acessar este módulo.",
		})
	}

	fmt.Println(userContext)

	userData := userContext.(utils.User)

	page := utils.ParseInt(context.DefaultQuery("page", "1"), context)
	pageSize := utils.ParseInt(context.DefaultQuery("page_size", "15"), context)

	startDate := utils.ParseDate(context.Query("start_date"), context)
	endDate := utils.ParseDate(context.Query("end_date"), context)

	payments, err := handler.service.GetPaymentSummary(
		utils.Pagination{
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

func (handler *Handler) GetAllPaymentHandler(context *gin.Context) {
	userContext, exist := context.Get("user")

	if !exist {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Este usuário não possui permissão para acessar este módulo.",
		})
	}

	userData := userContext.(utils.User)

	page := utils.ParseInt(context.DefaultQuery("page", "1"), context)
	pageSize := utils.ParseInt(context.DefaultQuery("page_size", "15"), context)

	status := utils.ParseInt(context.DefaultQuery("status", "1"), context)
	paymentType := utils.ParseInt(context.DefaultQuery("type", "1"), context)
	name := context.Query("name")

	startDate := utils.ParseDate(context.Query("start_date"), context)
	endDate := utils.ParseDate(context.Query("end_date"), context)

	installment := utils.ParseInt(context.DefaultQuery("installment", "0"), context)

	startPaymentDate := utils.ParseDate(context.Query("start_date"), context)
	endPaymentDate := utils.ParseDate(context.Query("end_date"), context)

	payments, err := handler.service.GetAllPaymentService(
		utils.Pagination{
			Page:     page,
			PageSize: pageSize,
		}, PaymentFilter{
			UserId:           userData.Id,
			Status:           status,
			Type:             paymentType,
			Name:             name,
			StartDate:        startDate,
			EndDate:          endDate,
			Installment:      installment,
			StartPaymentDate: startPaymentDate,
			EndPaymentDate:   endPaymentDate,
			Active:           true,
		},
	)
	if err != nil {
		log.Println(err)
		context.AbortWithError(http.StatusInternalServerError, err)
	}

	context.JSON(http.StatusOK, gin.H{
		"payments":  payments.Data,
		"page_info": payments.PageInfo,
	})
}

func (handler *Handler) GetPaymentHandler(context *gin.Context) {
	userContext, exist := context.Get("user")

	if !exist {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Este usuário não possui permissão para acessar este módulo.",
		})
	}

	userData := userContext.(utils.User)

	paymentId := utils.ParseInt(context.Param("paymentid"), context)

	payment, err := handler.service.GetPaymentByIdService(paymentId, userData.Id)

	if err != nil {
		log.Println(err)
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": payment,
	})

}

func (handler *Handler) SavePaymentHandler(context *gin.Context) {
	// paymentId := utils.ParseInt(context.Param("paymentid"), context)

}

func (handler *Handler) PayoffPaymentHandler(context *gin.Context) {
	// paymentId := utils.ParseInt(context.Param("paymentid"), context)

}
