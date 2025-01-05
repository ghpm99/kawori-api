package tag

type Tag struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	UserId int    `json:"-"`
}
