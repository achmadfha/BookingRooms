package authUsecase

import (
	"BookingRoom/model/dto"
	"BookingRoom/pkg/utils"
	"BookingRoom/src/auth"
	"errors"
	"fmt"
)

type authUC struct {
	authRepo auth.AuthRepository
}

func NewAuthUsecase(authRepo auth.AuthRepository) auth.AuthUsecase {
	return &authUC{authRepo}
}

func (e *authUC) Login(employees dto.LoginRequest) (token string, err error) {
	emp, err := e.authRepo.RetrieveEmployees(employees.Username)
	if err != nil {
		fmt.Println("Error Usecase > repo: ", err.Error())
		if err.Error() == "no rows" {
			return "", errors.New("01")
		}
		return "", err
	}

	if err = utils.VerifyPassword(emp.Password, employees.Password); err != nil {
		return "", errors.New("02")
	}

	token, err = utils.GenerateToken(emp.EmployeeId, string(emp.Position))
	if err != nil {
		return "", err
	}

	return token, err
}
