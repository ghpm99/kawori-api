package financial

type FinancialService interface {
	GetAllPayments() ([]*Payment, error)
}

func (s *FinancialService) GetAllPayments() ([]*Payment, error) {

	var array = []*Payment{{"12/12/24", 1, 1, 1, 1, 1, 1}}
	return array, nil
}
