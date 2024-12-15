package financial

import "time"

type PaymentSummaryFilter struct {
	UserId      int
	DataInicial time.Time
	DataFinal   time.Time
}

type Pagination struct {
	Page     int
	PageSize int
	HasNext  bool
	HasPrev  bool
}

type GetPaymentSummaryReturn struct {
	data     []Payment
	pageInfo Pagination
}
