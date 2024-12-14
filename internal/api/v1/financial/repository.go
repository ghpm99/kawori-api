package financial

import (
	"database/sql"
)

type Repository struct {
	dbContext *sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{database}
}

func (repository *Repository) GetAllPayments() ([]Payment, error) {
	data, err := repository.dbContext.Query("select * from financial_paymentsummary")
	if err != nil {
		return nil, err
	}
	var paymentsArray []Payment
	for data.Next() {
		var payment Payment

		if errPayment := data.Scan(
			&payment.PaymentsDate,
			&payment.UserId,
			&payment.Total,
			&payment.Debit,
			&payment.Credit,
			&payment.Dif,
			&payment.Accumulated,
		); errPayment != nil {
			return nil, errPayment

		}
		paymentsArray = append(paymentsArray, payment)
	}
	if errorSql := data.Err(); errorSql != nil {
		return nil, errorSql
	}

	return paymentsArray, nil
}
