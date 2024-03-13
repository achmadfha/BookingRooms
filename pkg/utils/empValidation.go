package utils

import (
	"BookingRoom/model/dto/employeesDto"
	"BookingRoom/model/dto/json"
	"regexp"
)

func ValidationEmployee(employee employeesDto.Employees) []json.ValidationField {
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

	phoneRegex := regexp.MustCompile(`^\d{9,13}$`)
	if !phoneRegex.MatchString(employee.PhoneNumber) {
		valError = append(valError, json.ValidationField{
			FieldName: "PhoneNumber",
			Message:   "is number.",
		})
	}

	return valError
}

func ValidationLogin(auth employeesDto.LoginRequest) []json.ValidationField {
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
