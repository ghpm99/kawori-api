package financial

type Payment struct {
	PaymentsDate string  `json:"payments_date"`
	UserId       int     `json:"user_id"`
	Total        int     `json:"total"`
	Debit        float64 `json:"debit"`
	Credit       float64 `json:"credit"`
	Dif          float64 `json:"dif"`
	Accumulated  float64 `json:"accumulated"`
}
