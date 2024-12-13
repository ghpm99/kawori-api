package financial

import (
	"fmt"
	"kawori/api/src/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payments struct {
	PaymentsDate string
	UserId       int
	Total        int
	Debit        int
	Credit       int
	Dif          int
	Accumulated  int
}

func FinancialEndpoint(context *gin.Context) {

	data, err := config.GlobalDatabase.Query("select * from financial_paymentsummary")
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, "Falhou em buscar pagamentos")
	} else {
		var paymentsArray []Payments

		for data.Next() {
			var payment Payments
			if errPayment := data.Scan(
				&payment.PaymentsDate,
				&payment.UserId,
				&payment.Total,
				&payment.Debit,
				&payment.Credit,
				&payment.Dif,
				&payment.Accumulated,
			); errPayment != nil {
				context.AbortWithStatusJSON(http.StatusInternalServerError, errPayment)
				break
			}
			paymentsArray = append(paymentsArray, payment)
		}
		if err = data.Err(); err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, "Falhou na execução da query")
		}

		context.JSON(http.StatusOK, paymentsArray)
	}
}
