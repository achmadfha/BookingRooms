package authRepository

import (
	"BookingRoom/model/dto/employeesDto"
	"BookingRoom/src/auth"
	"database/sql"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) auth.AuthRepository {
	return &authRepository{db}
}

func (e *authRepository) RetrieveEmployees(username string) (employeesDto.Employees, error) {
	var employee employeesDto.Employees
	err := e.db.QueryRow("SELECT employee_id, full_name, division, phone_number, position, username, password FROM employee WHERE username = $1", username).Scan(&employee.EmployeeId, &employee.FullName, &employee.Division, &employee.PhoneNumber, &employee.Position, &employee.Username, &employee.Password)
	if err != nil {
		return employeesDto.Employees{}, err
	}

	return employee, err
}

func (e *authRepository) RenewPassword(username, password string) error {
	query := "UPDATE employee SET password=$1 WHERE username=$1"
	_, err := e.db.Exec(query, password, username)
	return err
}
