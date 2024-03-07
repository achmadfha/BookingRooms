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

func (e *authUC) Login(employee employeesDto.LoginRequest) (token string, err error) {
	emp, err := e.authRepo.RetrieveEmployees(employee.Username)
	if err != nil {
		return "", errors.New("01")
	}

	if err = utils.VerifyPassword(emp.Password, employee.Password); err != nil {
		return "", errors.New("02")
	}

	token, err = utils.GenerateToken(emp.EmployeeId, string(emp.Position))
	if err != nil {
		return "", errors.New("03")
	}

	return token, err
}

func (e *authUC) UpdatePassword(employee employeesDto.PasswordRequest) error {
	emp, err := e.authRepo.RetrieveEmployees(employee.Username)
	if err != nil {
		return errors.New("01")
	}

	if err = utils.VerifyPassword(emp.Password, employee.OldPasswrod); err != nil {
		return errors.New("02")
	}

	if employee.NewPassword != employee.ConfirmPassword {
		return errors.New("03")
	}

	hashPass, err := utils.HashPassword(employee.NewPassword)
	if err != nil {
		return errors.New("04")
	}

	err = e.authRepo.RenewPassword(employee.Username, hashPass)
	if err != nil {
		return errors.New("05")
	}
	return err
}
