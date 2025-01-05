package payment

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kawori/api/pkg/database/queries"
)

type Repository struct {
	dbContext *sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{database}
}

func (repository *Repository) CreateTransaction(ctx context.Context) (*sql.Tx, error) {
	return repository.dbContext.BeginTx(ctx, nil)
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
			&payment.InvoiceId,
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
		&payment.InvoiceId,
	); err != nil {
		return Payment{}, err

	}

	return payment, nil
}

func (repository *Repository) CreatePayment(transaction *sql.Tx, payment Payment) (Payment, error) {

	fail := func(err error) (Payment, error) {
		return payment, fmt.Errorf("CreatePayment: %v", err)
	}

	defer transaction.Rollback()

	data, err := transaction.Exec(
		queries.CreatePayment,
		&payment.Type,
		&payment.Name,
		&payment.Date,
		&payment.Installments,
		&payment.PaymentDate,
		&payment.Fixed,
		&payment.Active,
		&payment.Value,
		&payment.Status,
		&payment.InvoiceId,
		&payment.UserId,
	)
	if err != nil {
		return fail(err)
	}
	paymentId, err := data.LastInsertId()

	if err != nil {
		return fail(err)
	}

	payment.Id = int(paymentId)

	return payment, nil
}

func (repository *Repository) UpdatePayment(transaction *sql.Tx, payment Payment) (bool, error) {

	fail := func(err error) (bool, error) {
		return false, fmt.Errorf("UpdatePayment: %v", err)
	}

	defer transaction.Rollback()

	result, err := transaction.Exec(
		queries.UpdatePayment,
		&payment.Id,
		&payment.UserId,
		&payment.Type,
		&payment.Name,
		&payment.Date,
		&payment.Installments,
		&payment.PaymentDate,
		&payment.Fixed,
		&payment.Active,
		&payment.Value,
		&payment.Status,
		&payment.InvoiceId,
	)

	if err != nil {
		return fail(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return fail(err)
	}

	if rowsAffected > 1 {
		transaction.Rollback()
		return fail(errors.New("multiple rows affected"))
	}

	return rowsAffected == 1, nil
}
