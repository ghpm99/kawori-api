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
	Status           utils.Optional[int]       `filter:"status"`
	Type             utils.Optional[int]       `filter:"type"`
	Name             utils.Optional[string]    `filter:"name"`
	StartDate        time.Time                 `filter:"start_date"`
	EndDate          time.Time                 `filter:"end_date"`
	Installment      utils.Optional[int]       `filter:"installment"`
	StartPaymentDate utils.Optional[time.Time] `filter:"start_payment_date"`
	EndPaymentDate   utils.Optional[time.Time] `filter:"end_payment_date"`
	Active           bool                      `filter:"active"`
	Fixed            utils.Optional[bool]      `filter:"fixed"`
}

type GetPaymentSummaryReturn struct {
	data     []PaymentSummary
	pageInfo utils.Pagination
}

type GetPaymentReturn struct {
	Data     []Payment        `json:"data"`
	PageInfo utils.Pagination `json:"page_info"`
}
