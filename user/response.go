package user

type UserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  int    `json:"phone"`
	Status string `json:"status"`
}
