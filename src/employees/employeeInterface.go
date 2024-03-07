package employees

import "BookingRoom/model/dto/employeesDto"

type EmployeeRepository interface {
	RetrieveEmployee() ([]employeesDto.Employees, error)
	RetrieveEmployeeById(id string) (employeesDto.Employees, error)
	CreateEmployees(employee *employeesDto.Employees) error
	RenewEmployee(employee employeesDto.Employees) error
	RemoveEmployeeById(id string) error
	CountEmployees(page, size int) (int, error)
}

type EmployeeUsecase interface {
	GetEmployee(page, size string) (employee []employeesDto.Employees, pagination interface{}, err error)
	GetEmployeeById(id string) (employeesDto.Employees, error)
	StoreEmployee(employee *employeesDto.Employees) error
	UpdateEmployee(employee employeesDto.Employees) error
	DeleteEmployeeById(id string) error
}
