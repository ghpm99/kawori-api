package payment

import (
	"database/sql"
	"errors"
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

func (repository *Repository) GetAllPayments(pagination Pagination, filters PaymentFilter) (GetPaymentReturn, error) {
	data, err := repository.dbContext.Query(
		queries.GetAllPayments,
		filters.Status,
		filters.Type,
		filters.Name,
		filters.StartDate,
		filters.EndDate,
		filters.installment,
		filters.StartPaymentDate,
		filters.EndPaymentDate,
		filters.Fixed,
		filters.Active,
		filters.UserId,
		pagination.PageSize,
		pagination.Page,
	)

	if err != nil {
		return GetPaymentReturn{}, err
	}
	var paymentsArray []Payment
	for data.Next() {
		var payment Payment

		if errPayment := data.Scan(
			&payment.Id,
			&payment.Status,
			&payment.Type,
			&payment.Name,
			&payment.Date,
			&payment.Installments,
			&payment.PaymentDate,
			&payment.Fixed,
			&payment.Value,
			&payment.Invoice,
		); errPayment != nil {
			return GetPaymentReturn{}, errPayment

		}
		paymentsArray = append(paymentsArray, payment)
	}
	if errorSql := data.Err(); errorSql != nil {
		return GetPaymentReturn{}, errorSql
	}

	if pagination.Page > 1 {
		pagination.HasPrev = true
	}

	if paymentsArray == nil {
		paymentsArray = []Payment{}
	}

	return GetPaymentReturn{
		data:     paymentsArray,
		pageInfo: pagination,
	}, nil
}

func (repository *Repository) GetPaymentById(idPayment int, IdUser int) (Payment, error) {
	data, err := repository.dbContext.Query(
		queries.GetPayment,
		idPayment,
		IdUser,
	)
	if err != nil {
		return Payment{}, err
	}
	var payment Payment

	hasData := data.Next()

	if !hasData {
		return Payment{}, errors.New("no data")
	}
	if err := data.Scan(
		&payment.Id,
		&payment.Status,
		&payment.Type,
		&payment.Name,
		&payment.Date,
		&payment.Installments,
		&payment.PaymentDate,
		&payment.Fixed,
		&payment.Active,
		&payment.Value,
		&payment.Invoice,
	); err != nil {
		return Payment{}, err

	}

	return payment, nil
}

func (repository *Repository) CreatePayment(payment Payment) (Payment, error){
	data, err := repository.dbContext.
}