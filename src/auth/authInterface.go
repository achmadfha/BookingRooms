package auth

import "BookingRoom/model/dto/employeesDto"

type AuthRepository interface {
	RetrieveEmployees(username string) (employeesDto.Employees, error)
	RenewPassword(username, password string) error
}

type AuthUsecase interface {
	Login(employee employeesDto.LoginRequest) (token string, err error)
	UpdatePassword(employee employeesDto.Employees) error
}
