package payment

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

type Payment struct {
	Id           int       `json:"id"`
	Status       int       `json:"status"`
	Type         int       `json:"type"`
	Name         string    `json:"name"`
	Date         time.Time `json:"date"`
	Installments int       `json:"installments"`
	PaymentDate  time.Time `json:"payment_date"`
	Fixed        bool      `json:"fixed"`
	Active       bool      `json:"active"`
	Value        float64   `json:"value"`
	InvoiceId    int       `json:"invoice"`
	UserId       int       `json:"-"`
}
