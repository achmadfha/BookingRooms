package employeeUsecase

import (
	"BookingRoom/model/dto"
	"BookingRoom/src/employees"
)

type employeeUC struct {
	employeeRepo employees.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo employees.EmployeeRepository) employees.EmployeeUsecase {
	return &employeeUC{employeeRepo}
}

func (e *employeeUC) GetLogin(username string) (dto.Employees, error) {
	employees, err := e.employeeRepo.GetLogin(username)
	if err != nil {
		return dto.Employees{}, err
	}

	return employees, err
}
