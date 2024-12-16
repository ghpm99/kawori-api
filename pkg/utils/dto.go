package utils

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsStaff     bool   `json:"is_staff"`
	IsActive    bool   `json:"is_active"`
	IsSuperuser bool   `json:"is_superuser"`
	LastLogin   string `json:"last_login"`
	DateJoined  string `json:"date_joined"`
}
