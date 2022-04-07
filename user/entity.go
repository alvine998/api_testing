package user

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Phone     int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
