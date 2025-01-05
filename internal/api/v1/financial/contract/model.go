package contract

type Contract struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	ValueClosed float64 `json:"value_closed"`
	ValueOpen   float64 `json:"value_open"`
	UserId      int     `json:"-"`
}
