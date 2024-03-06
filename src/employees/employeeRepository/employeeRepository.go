package employeeRepository

import (
	"BookingRoom/model/dto"
	"BookingRoom/pkg/utils"
	"BookingRoom/src/employees"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) employees.EmployeeRepository {
	return &employeeRepository{db}
}

func (e *employeeRepository) RetrieveEmployee() ([]dto.Employees, error) {
	// var rows *sql.Rows
	// var err error
	query := "SELECT employee_id, full_name, division, phone_number, position FROM employee"
	rows, err := e.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []dto.Employees
	for rows.Next() {
		var employee dto.Employees
		err := rows.Scan(&employee.EmployeeId, &employee.FullName, &employee.Division, &employee.PhoneNumber, &employee.Position)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, err
}

func (e *employeeRepository) RetrieveEmployeeById(id string) (dto.Employees, error) {
	var employee dto.Employees
	row := e.db.QueryRow("SELECT employee_id, full_name, division, phone_number, position FROM employee WHERE employee_id=$1", id)
	err := row.Scan(&employee.EmployeeId, &employee.FullName, &employee.Division, &employee.PhoneNumber, &employee.Position)
	if err != nil {
		return dto.Employees{}, errors.New("1")
	}

	return employee, err
}

func (e *employeeRepository) CreateEmployees(employee *dto.Employees) error {
	password, err := utils.HashPassword(employee.Username)
	if err != nil {
		return err
	}

	employee.EmployeeId = uuid.New()
	query := "INSERT INTO employee (employee_id, full_name, division, phone_number, position, username, password) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = e.db.Exec(query, employee.EmployeeId, employee.FullName, employee.Division, employee.PhoneNumber, employee.Position, employee.Username, password)

	return err
}

func (e *employeeRepository) RenewEmployee(employee dto.Employees) error {
	query := "UPDATE employee SET full_name=$1, division=$2, phone_number=$3, position=$4 WHERE employee_id=$5"
	_, err := e.db.Exec(query, employee.FullName, employee.Division, employee.PhoneNumber, employee.Position, employee.EmployeeId)
	return err
}

func (e *employeeRepository) RemoveEmployeeById(id string) error {
	query := "DELETE FROM employee WHERE employee_id=$1"
	_, err := e.db.Exec(query, id)
	return err
}
