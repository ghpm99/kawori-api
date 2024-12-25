package payment

import "time"

type PaymentSummaryFilter struct {
	UserId    int
	StartDate time.Time
	EndDate   time.Time
}

type PaymentFilter struct {
	UserId           int
	Status           int
	Type             int
	Name             string
	StartDate        time.Time
	EndDate          time.Time
	installments     int
	StartPaymentDate time.Time
	EndPaymentDate   time.Time
	Fixed            bool
	Active           bool
}

type Pagination struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	HasNext  bool `json:"has_next"`
	HasPrev  bool `json:"has_prev"`
}

type GetPaymentSummaryReturn struct {
	data     []PaymentSummary
	pageInfo Pagination
}

type GetPaymentReturn struct {
	data     []Payment
	pageInfo Pagination
}
