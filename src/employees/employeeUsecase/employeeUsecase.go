package employeeUsecase

import (
	"BookingRoom/model/dto"
	"BookingRoom/pkg/utils"
	"BookingRoom/src/employees"
	"errors"
	"fmt"
)

type employeeUC struct {
	employeeRepo employees.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo employees.EmployeeRepository) employees.EmployeeUsecase {
	return &employeeUC{employeeRepo}
}

func (e *employeeUC) Login(employees dto.LoginRequest) (token string, err error) {
	emp, err := e.employeeRepo.RetrieveEmployees(employees.Username)
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
