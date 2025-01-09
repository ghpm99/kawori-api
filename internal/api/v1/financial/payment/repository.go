package payment

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"kawori/api/pkg/database/queries"
	"kawori/api/pkg/utils"
)

type Repository struct {
	dbContext *sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{database}
}

func (repository *Repository) GetPaymentSummary(pagination utils.Pagination, filters PaymentSummaryFilter) (GetPaymentSummaryReturn, error) {

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

func (repository *Repository) GetAllPayments(pagination utils.Pagination, filters PaymentFilter) (GetPaymentReturn, error) {

	if json, _ := json.Marshal(filters); json != nil {
		fmt.Println(string(json))

	}

	args := []interface{}{
		pagination.PageSize,
		pagination.Page,
		filters.UserId,
		filters.Status.Value,
		filters.Type.Value,
		fmt.Sprintf("%%%s%%", filters.Name.Value),
		filters.StartDate,
		filters.EndDate,
		filters.Installment.Value,
		filters.StartPaymentDate.Value,
		filters.EndPaymentDate.Value,
		filters.Active,
		!filters.Fixed.HasValue,
		filters.Fixed.Value,
	}

	query := queries.GetAllPayments

	utils.PrintQuery(query, args)

	data, err := repository.dbContext.Query(
		query,
		args...,
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
		Data:     paymentsArray,
		PageInfo: pagination,
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

	args := []interface{}{
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
	}

	query := queries.CreatePayment

	utils.PrintQuery(query, args)

	data, err := transaction.Exec(
		query,
		args...,
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
