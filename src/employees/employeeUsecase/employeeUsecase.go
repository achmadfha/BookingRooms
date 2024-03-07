package employeeUsecase

import (
	"BookingRoom/model/dto/employeesDto"
	"BookingRoom/model/dto/json"
	"BookingRoom/src/employees"
	"math"
	"strconv"
)

type employeeUC struct {
	employeeRepo employees.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo employees.EmployeeRepository) employees.EmployeeUsecase {
	return &employeeUC{employeeRepo}
}

func (e *employeeUC) GetEmployee(page, size string) (employee []employeesDto.Employees, pagination interface{}, err error) {
	employee, err = e.employeeRepo.RetrieveEmployee()
	if err != nil {
		return nil, nil, err
	}

	var pageInt, sizeInt int
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil || pageInt < 1 {
			return
		}
	} else {
		pageInt = 1
	}

	if size != "" {
		sizeInt, err = strconv.Atoi(size)
		if err != nil {
			return
		}
	} else {
		sizeInt = 5
	}

	totalData, err := e.employeeRepo.CountEmployees(pageInt, sizeInt)
	if err != nil {
		return nil, nil, err
	}

	totalPages := int(math.Ceil(float64(totalData) / float64(sizeInt)))
	if pageInt > totalPages {
		return nil, json.Pagination{}, err
	}

	if totalPages == 0 && totalData > 0 {
		totalPages = 1
	}

	pagination = json.Pagination{
		CurrentPage:  pageInt,
		TotalPages:   totalPages,
		TotalRecords: totalData,
	}

	return employee, pagination, nil
}

func (e *employeeUC) GetEmployeeById(id string) (employeesDto.Employees, error) {
	employee, err := e.employeeRepo.RetrieveEmployeeById(id)
	if err != nil {
		return employeesDto.Employees{}, err
	}

	return employee, nil
}

func (e *employeeUC) StoreEmployee(employee *employeesDto.Employees) error {
	err := e.employeeRepo.CreateEmployees(employee)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeUC) UpdateEmployee(employee employeesDto.Employees) error {

	// Validasi
	_, err := e.employeeRepo.RetrieveEmployeeById(employee.EmployeeId.String())
	if err != nil {
		return err
	}

	err = e.employeeRepo.RenewEmployee(employee)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeUC) DeleteEmployeeById(id string) error {
	// Validasi id jika diperlukan

	_, err := e.employeeRepo.RetrieveEmployeeById(id)
	if err != nil {
		return err
	}

	err = e.employeeRepo.RemoveEmployeeById(id)
	if err != nil {
		return err
	}

	return nil
}
