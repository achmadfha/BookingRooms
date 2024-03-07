package employeesDto

import (
	"BookingRoom/model/dto/json"

	"github.com/google/uuid"
)

type (
	Employees struct {
		EmployeeId  uuid.UUID `json:"employeeId"`
		FullName    string    `json:"fullName"`
		Division    string    `json:"division"`
		PhoneNumber string    `json:"phoneNumber"`
		Position    string    `json:"position"`
		Username    string    `json:"username,omitempty"`
		Password    string    `json:"password,omitempty"`
	}

	EmployeeResponse struct {
		EmployeeId  uuid.UUID `json:"employeeId"`
		FullName    string    `json:"fullName"`
		Division    string    `json:"division"`
		PhoneNumber string    `json:"phoneNumber"`
		Position    string    `json:"position"`
		Username    string    `json:"username"`
	}

	EmployeeRequest struct {
		FullName    string `json:"fullName"`
		Division    string `json:"division"`
		PhoneNumber string `json:"phoneNumber"`
		Position    string `json:"position"`
		Username    string `json:"username"`
	}

	LoginRequest struct {
		Username string `json:"username" bidding:"required"`
		Password string `json:"password" bidding:"required"`
	}
)

func ValidationEmployee(employee Employees) []json.ValidationField {
	var valError []json.ValidationField
	if employee.FullName == "" {
		valError = append(valError, json.ValidationField{
			FieldName: "FullName",
			Message:   "is required.",
		})
	}
	if employee.Division == "" {
		valError = append(valError, json.ValidationField{
			FieldName: "Division",
			Message:   "is required.",
		})
	}
	if employee.PhoneNumber == "" {
		valError = append(valError, json.ValidationField{
			FieldName: "PhoneNumber",
			Message:   "is required.",
		})
	}
	if employee.Position == "" {
		valError = append(valError, json.ValidationField{
			FieldName: "Position",
			Message:   "is required.",
		})
	}
	if employee.Username == "" {
		valError = append(valError, json.ValidationField{
			FieldName: "Username",
			Message:   "is required.",
		})
	}

	return valError
}

func ValidationLogin(auth LoginRequest) []json.ValidationField {
	var valError []json.ValidationField
	if auth.Username == "" {
		valError = append(valError, json.ValidationField{
			FieldName: "Username",
			Message:   "is required.",
		})
	}
	if auth.Password == "" {
		valError = append(valError, json.ValidationField{
			FieldName: "Password",
			Message:   "is required.",
		})
	}
	return valError
}
