package payment

import "time"

type PaymentSummaryFilter struct {
	UserId    int
	StartDate time.Time
	EndDate   time.Time
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
