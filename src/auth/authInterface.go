package auth

import "BookingRoom/model/dto"

type AuthRepository interface {
	RetrieveEmployees(username string) (dto.Employees, error)
}

type AuthUsecase interface {
	Login(employee dto.LoginRequest) (token string, err error)
}
