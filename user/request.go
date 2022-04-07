package user

import "encoding/json"

type UserRequest struct {
	Name   string      `json:"name" binding:"required"`
	Email  string      `json:"email" binding:"required"`
	Phone  json.Number `json:"phone" binding:"required,number"`
	Status string      `json:"status" binding:"required"`
}
