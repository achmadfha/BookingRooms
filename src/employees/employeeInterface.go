package employees

import "BookingRoom/model/dto"

type EmployeeRepository interface {
	GetLogin(username string) (dto.Employees, error)
}

type EmployeeUsecase interface {
	GetLogin(username string) (dto.Employees, error)
}
