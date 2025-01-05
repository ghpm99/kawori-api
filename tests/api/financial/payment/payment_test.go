package payment_test

import (
	"encoding/json"
	"fmt"
	"kawori/api/internal/api/v1/financial/payment"
	"kawori/api/internal/app"
	"kawori/api/pkg/utils"
	"kawori/api/tests"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPaymentSummary(t *testing.T) {

	authMiddleware := func(c *gin.Context) {
		var user = utils.User{
			Id:          1,
			Name:        "teste",
			Username:    "teste",
			FirstName:   "teste",
			LastName:    "teste",
			Email:       "teste",
			IsStaff:     false,
			IsActive:    true,
			IsSuperuser: false,
			LastLogin:   "",
			DateJoined:  "",
		}
		c.Set("user", user)
		c.Next()
	}

	database := tests.ConfigInMemoryDatabase()
	defer database.Close()
	CreatePaymentTableFixture(database)

	repository := payment.NewRepository(database)
	service := payment.NewService(repository)
	handler := payment.NewHandler(service)

	CreatePaymentsDataFixture(repository, database)

	router := app.SetUpRouter()
	router.Use(authMiddleware)

	router.GET("/", handler.GetAllPaymentHandler)

	req, _ := http.NewRequest(http.MethodGet, "/?start_date=2024-01-01&end_date=2025-01-01&name=teste", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	fmt.Println(w.Body)
	dataRef, err := time.Parse("2006-01-02", "2024-01-02")
	if err != nil {
		log.Fatalf("Falha ao popular tabela: %v", err)
	}

	var paymentArray = []payment.Payment{
		{
			Id:           1,
			Status:       1,
			Type:         1,
			Name:         "teste",
			Date:         dataRef,
			Installments: 1,
			PaymentDate:  dataRef,
			Fixed:        false,
			Active:       true,
			Value:        100.0,
			InvoiceId:    1,
			UserId:       1,
		},
	}

	var pagination = utils.Pagination{
		Page:     1,
		PageSize: 15,
		HasNext:  false,
		HasPrev:  false,
	}

	expectedReturn := payment.GetPaymentReturn{
		Data:     paymentArray,
		PageInfo: pagination,
	}
	paymentJson, _ := json.Marshal(expectedReturn)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(paymentJson), w.Body.String())
}
