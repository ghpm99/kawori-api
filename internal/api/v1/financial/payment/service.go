package payment

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) GetPaymentSummary(pagination Pagination, filters PaymentSummaryFilter) (GetPaymentSummaryReturn, error) {
	return service.repository.GetPaymentSummary(pagination, filters)
}

func (service *Service) GetAllPaymentService(pagination Pagination, filters PaymentFilter) (GetPaymentReturn, error) {
	return service.GetAllPaymentService(pagination, filters)
}
