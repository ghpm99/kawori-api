package financial

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) GetAllPayments() ([]Payment, error) {
	return service.repository.GetAllPayments()
}
