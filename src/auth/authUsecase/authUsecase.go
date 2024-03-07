package authUsecase

import (
	"BookingRoom/model/dto/employeesDto"
	"BookingRoom/pkg/utils"
	"BookingRoom/src/auth"
	"errors"
)

type authUC struct {
	authRepo auth.AuthRepository
}

func NewAuthUsecase(authRepo auth.AuthRepository) auth.AuthUsecase {
	return &authUC{authRepo}
}

func (e *authUC) Login(employees employeesDto.LoginRequest) (token string, err error) {
	emp, err := e.authRepo.RetrieveEmployees(employees.Username)
	if err != nil {
		return "", errors.New("01")
	}

	if err = utils.VerifyPassword(emp.Password, employees.Password); err != nil {
		return "", errors.New("02")
	}

	token, err = utils.GenerateToken(emp.EmployeeId, string(emp.Position))
	if err != nil {
		return "", errors.New("03")
	}

	return token, err
}

func (e *authUC) UpdatePassword(employee employeesDto.Employees) error {
	return nil
}
