package employeeRepository

import (
	"BookingRoom/model/dto"
	"BookingRoom/src/employees"
	"database/sql"
)

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) employees.EmployeeRepository {
	return &employeeRepository{db}
}

func (e *employeeRepository) GetLogin(username string) (dto.Employees, error) {
	var employee dto.Employees
	err := e.db.QueryRow("SELECT employee_id, full_name, division, phone_number, position, username, password FROM employee WHERE username = $1", username).Scan(&employee.EmployeeId, &employee.FullName, &employee.Division, &employee.PhoneNumber, &employee.Position, &employee.Username, &employee.Password)
	if err != nil {
		return dto.Employees{}, err
	}

	return employee, err
}
