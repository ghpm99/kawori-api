package payment

import (
	"context"
	"kawori/api/pkg/utils"
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
