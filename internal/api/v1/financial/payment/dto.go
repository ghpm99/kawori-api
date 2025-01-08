package payment

import (
	"kawori/api/pkg/utils"
	"time"
)

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
	Installment      int
	StartPaymentDate time.Time
	EndPaymentDate   time.Time
	Active           bool
}

type GetPaymentSummaryReturn struct {
	data     []PaymentSummary
	pageInfo utils.Pagination
}

type GetPaymentReturn struct {
	Data     []Payment        `json:"data"`
	PageInfo utils.Pagination `json:"page_info"`
}
