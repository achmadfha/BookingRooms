package employees

import "BookingRoom/model/dto"

type EmployeeRepository interface {
	RetrieveEmployee() ([]dto.Employees, error)
	RetrieveEmployeeById(id string) (dto.Employees, error)
	CreateEmployees(employee *dto.Employees) error
	RenewEmployee(employee dto.Employees) error
	RemoveEmployeeById(id string) error
}

type EmployeeUsecase interface {
	GetEmployee() ([]dto.Employees, error)
	GetEmployeeById(id string) (dto.Employees, error)
	StoreEmployee(employee *dto.Employees) error
	UpdateEmployee(employee dto.Employees) error
	DeleteEmployeeById(id string) error
}
