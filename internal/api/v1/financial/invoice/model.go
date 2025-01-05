package invoice

import "time"

type Invoice struct {
	Id           int       `json:"id"`
	Status       int       `json:"status"`
	Name         string    `json:"name"`
	Installments int       `json:"installments"`
	Value        float64   `json:"value"`
	Date         time.Time `json:"date"`
	ContractId   int       `json:"contract_id"`
	Active       bool      `json:"active"`
	Fixed        bool      `json:"fixed"`
	PaymentDate  time.Time `json:"payments_date"`
	Type         int       `json:"type"`
	ValueClosed  float64   `json:"value_closed"`
	ValueOpen    float64   `json:"value_open"`
	UserId       int       `json:"-"`
}
