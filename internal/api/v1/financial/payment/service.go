package payment

import (
	"context"
	"database/sql"
	"kawori/api/pkg/utils"
	"time"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) GetPaymentSummary(pagination utils.Pagination, filters PaymentSummaryFilter) (GetPaymentSummaryReturn, error) {
	return service.repository.GetPaymentSummary(pagination, filters)
}

func (service *Service) GetAllPaymentService(pagination utils.Pagination, filters PaymentFilter) (GetPaymentReturn, error) {
	ctx := context.Background()
	transaction, _ := service.repository.dbContext.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})

	defer transaction.Rollback()

	dataRef, _ := time.Parse("2006-01-02", "2024-01-02")

	service.repository.CreatePayment(transaction, Payment{
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
	})

	transaction.Commit()

	return service.repository.GetAllPayments(pagination, filters)
}

func (service *Service) GetPaymentByIdService(idPayment int, IdUser int) (Payment, error) {
	return service.repository.GetPaymentById(idPayment, IdUser)
}

func (service *Service) UpdatePaymentService(payment Payment) (bool, error) {
	ctx := context.Background()
	transaction, err := service.repository.dbContext.BeginTx(ctx, nil)

	if err != nil {
		return false, err
	}
	result, err := service.repository.UpdatePayment(transaction, payment)

	if result {
		err = transaction.Commit()
	}

	return result, err
}
