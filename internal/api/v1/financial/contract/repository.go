package contract

import (
	"database/sql"
	"kawori/api/pkg/database/queries"
)

type Repository struct {
	dbContext *sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{database}
}

func (repository *Repository) GetPaymentSummary(pagination Pagination, filters PaymentSummaryFilter) (GetPaymentSummaryReturn, error) {

	data, err := repository.dbContext.Query(
		queries.GetPaymentSummary,
		filters.UserId,
		filters.StartDate,
		filters.EndDate,
		pagination.PageSize,
		pagination.Page,
	)

	if err != nil {
		return GetPaymentSummaryReturn{}, err
	}
	var paymentsArray []PaymentSummary
	for data.Next() {
		var payment PaymentSummary

		if errPayment := data.Scan(
			&payment.PaymentsDate,
			&payment.UserId,
			&payment.Total,
			&payment.Debit,
			&payment.Credit,
			&payment.Dif,
			&payment.Accumulated,
		); errPayment != nil {
			return GetPaymentSummaryReturn{}, errPayment

		}
		paymentsArray = append(paymentsArray, payment)
	}
	if errorSql := data.Err(); errorSql != nil {
		return GetPaymentSummaryReturn{}, errorSql
	}

	if pagination.Page > 1 {
		pagination.HasPrev = true
	}

	if paymentsArray == nil {
		paymentsArray = []PaymentSummary{}
	}

	return GetPaymentSummaryReturn{
		data:     paymentsArray,
		pageInfo: pagination,
	}, nil
}
