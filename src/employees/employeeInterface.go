package employees

import "BookingRoom/model/dto"

type EmployeeRepository interface {
	RetrieveEmployees(username string) (dto.Employees, error)
}

type EmployeeUsecase interface {
	Login(employee dto.LoginRequest) (token string, err error)
}
