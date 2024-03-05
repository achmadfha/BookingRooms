package dto

import "github.com/google/uuid"

type Employees struct {
	EmployeeId  uuid.UUID `json:"employeeId"`
	FullName    string    `json:"fullName"`
	Division    string    `json:"division"`
	PhoneNumber string    `json:"phoneNumber"`
	Position    string    `json:"position"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username" bidding:"required"`
	Password string `json:"password" bidding:"required"`
}
