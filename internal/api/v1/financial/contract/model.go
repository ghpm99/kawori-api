package contract

import "time"

type PaymentSummary struct {
	PaymentsDate time.Time `json:"payments_date"`
	UserId       int       `json:"-"`
	Total        int       `json:"total"`
	Debit        float64   `json:"debit"`
	Credit       float64   `json:"credit"`
	Dif          float64   `json:"dif"`
	Accumulated  float64   `json:"accumulated"`
}
