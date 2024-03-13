package employeeRepository

import (
	"BookingRoom/model/dto/employeesDto"
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

func (e *employeeRepository) RetrieveEmployee() ([]employeesDto.Employees, error) {
	// var rows *sql.Rows
	// var err error
	query := "SELECT employee_id, full_name, division, phone_number, position FROM employee"
	rows, err := e.db.Query(query)
	if err != nil {
		return nil, errors.New("1")
	}
	defer rows.Close()

	var employees []employeesDto.Employees
	for rows.Next() {
		var employee employeesDto.Employees
		err := rows.Scan(&employee.EmployeeId, &employee.FullName, &employee.Division, &employee.PhoneNumber, &employee.Position)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, err
}

func (e *employeeRepository) RetrieveEmployeeById(id string) (employeesDto.Employees, error) {
	var employee employeesDto.Employees
	row := e.db.QueryRow("SELECT employee_id, full_name, division, phone_number, position FROM employee WHERE employee_id=$1", id)
	err := row.Scan(&employee.EmployeeId, &employee.FullName, &employee.Division, &employee.PhoneNumber, &employee.Position)
	if err != nil {
		if err == sql.ErrNoRows {
			return employeesDto.Employees{}, errors.New("02")
		}
		return employeesDto.Employees{}, errors.New("1")
	}

	return employee, err
}

func (e *employeeRepository) CreateEmployees(employee *employeesDto.Employees) error {
	password, err := utils.HashPassword(employee.Username)
	if err != nil {
		return err
	}

	employee.EmployeeId = uuid.New()
	query := "INSERT INTO employee (employee_id, full_name, division, phone_number, position, username, password) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err = e.db.Exec(query, employee.EmployeeId, employee.FullName, employee.Division, employee.PhoneNumber, employee.Position, employee.Username, password)

	return err
}

func (e *employeeRepository) RenewEmployee(employee employeesDto.Employees) error {
	query := "UPDATE employee SET full_name=$1, division=$2, phone_number=$3, position=$4 WHERE employee_id=$5"
	_, err := e.db.Exec(query, employee.FullName, employee.Division, employee.PhoneNumber, employee.Position, employee.EmployeeId)
	return err
}

func (e *employeeRepository) RemoveEmployeeById(id string) error {
	query := "DELETE FROM employee WHERE employee_id=$1"
	_, err := e.db.Exec(query, id)
	return err
}

func (e *employeeRepository) CountEmployees(page, size int) (int, error) {
	var totalData int
	offset := (page - 1) * size
	query := "SELECT COUNT(*) FROM employee LIMIT $1 OFFSET $2"
	err := e.db.QueryRow(query, size, offset).Scan(&totalData)
	if err != nil {
		return 0, err
	}

	return totalData, nil
}
