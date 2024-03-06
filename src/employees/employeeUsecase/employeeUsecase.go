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

func (e *employeeUC) GetEmployee() ([]dto.Employees, error) {
	employee, err := e.employeeRepo.RetrieveEmployee()
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (e *employeeUC) GetEmployeeById(id string) (dto.Employees, error) {
	employee, err := e.employeeRepo.RetrieveEmployeeById(id)
	if err != nil {
		return dto.Employees{}, err
	}

	return employee, nil
}

func (e *employeeUC) StoreEmployee(employee *dto.Employees) error {
	err := e.employeeRepo.CreateEmployees(employee)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeUC) UpdateEmployee(employee dto.Employees) error {

	// Validasi

	err := e.employeeRepo.RenewEmployee(employee)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeUC) DeleteEmployeeById(id string) error {
	// Validasi id jika diperlukan

	err := e.employeeRepo.RemoveEmployeeById(id)
	if err != nil {
		return err
	}

	return nil
}
