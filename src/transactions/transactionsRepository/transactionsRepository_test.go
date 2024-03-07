package transactionsRepository

import (
	"BookingRoom/model/dto/transactionsDto"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestRetrieveAllTransactions_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsRepository(db)

	// Mock data
	createdAt := time.Now()
	updatedAt := time.Now()
	expectedTransactions := []transactionsDto.Transactions{
		{
			ID:          uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
			EmployeeId:  uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c9"),
			RoomId:      uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c0"),
			StartDate:   "2024-03-01",
			EndDate:     "2024-03-05",
			Description: "Test Transaction 1",
			Status:      "completed",
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		},
		{
			ID:          uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"),
			EmployeeId:  uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c9"),
			RoomId:      uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c0"),
			StartDate:   "2024-03-02",
			EndDate:     "2024-03-06",
			Description: "Test Transaction 2",
			Status:      "pending",
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		},
	}

	// Mock the query and expected result
	page := 1
	pageSize := 10
	startDate := "2024-03-01"
	endDate := "2024-03-10"

	rows := sqlmock.NewRows([]string{"transaction_id", "employee_id", "room_id", "start_date", "end_date", "description", "status", "created_at", "updated_at"}).
		AddRow(expectedTransactions[0].ID, expectedTransactions[0].EmployeeId, expectedTransactions[0].RoomId, expectedTransactions[0].StartDate, expectedTransactions[0].EndDate, expectedTransactions[0].Description, expectedTransactions[0].Status, expectedTransactions[0].CreatedAt, expectedTransactions[0].UpdatedAt).
		AddRow(expectedTransactions[1].ID, expectedTransactions[1].EmployeeId, expectedTransactions[1].RoomId, expectedTransactions[1].StartDate, expectedTransactions[1].EndDate, expectedTransactions[1].Description, expectedTransactions[1].Status, expectedTransactions[1].CreatedAt, expectedTransactions[1].UpdatedAt)
	mock.ExpectQuery("SELECT transaction_id, employee_id, room_id, start_date, end_date, description, status, created_at, updated_at FROM transactions WHERE created_at BETWEEN").
		WithArgs(startDate, endDate, pageSize, (page-1)*pageSize).
		WillReturnRows(rows)

	// Call the function under test
	transactions, err := repo.RetrieveAllTransactions(page, pageSize, startDate, endDate)

	// Check for errors
	assert.NoError(t, err)

	// Check if the transactions match the expected result
	assert.True(t, reflect.DeepEqual(expectedTransactions, transactions))

	// Verify that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRetrieveAllTransactions_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsRepository(db)

	// Mock the query to return an error
	page := 1
	pageSize := 10
	startDate := "2024-03-01"
	endDate := "2024-03-10"

	mock.ExpectQuery("SELECT transaction_id, employee_id, room_id, start_date, end_date, description, status, created_at, updated_at FROM transactions WHERE created_at BETWEEN").
		WithArgs(startDate, endDate, pageSize, (page-1)*pageSize).
		WillReturnError(errors.New("database error"))

	// Call the function under test
	transactions, err := repo.RetrieveAllTransactions(page, pageSize, startDate, endDate)

	// Check for errors
	assert.Error(t, err)
	assert.Nil(t, transactions)

	// Verify that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRetrieveAllTransactions_ErrorScanningRow(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsRepository(db)

	// Mock the query to return rows with an error during scanning
	page := 1
	pageSize := 10
	startDate := "2024-03-01"
	endDate := "2024-03-10"

	rows := sqlmock.NewRows([]string{"transaction_id", "employee_id", "room_id", "start_date", "end_date", "description", "status", "created_at", "updated_at"}).
		AddRow("invalid", "invalid", "invalid", "invalid", "invalid", "invalid", "invalid", "invalid", "invalid")

	mock.ExpectQuery("SELECT transaction_id, employee_id, room_id, start_date, end_date, description, status, created_at, updated_at FROM transactions WHERE created_at BETWEEN").
		WithArgs(startDate, endDate, pageSize, (page-1)*pageSize).
		WillReturnRows(rows).
		WillReturnError(errors.New("scanning error"))

	// Call the function under test
	transactions, err := repo.RetrieveAllTransactions(page, pageSize, startDate, endDate)

	// Check for errors
	assert.Error(t, err)
	assert.Nil(t, transactions)

	// Verify that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRetrieveAllTransactions_ErrorIteratingRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTransactionsRepository(db)

	// Mock the query to return an error during rows.Err()
	page := 1
	pageSize := 10
	startDate := "2024-03-01"
	endDate := "2024-03-10"

	mock.ExpectQuery("SELECT transaction_id, employee_id, room_id, start_date, end_date, description, status, created_at, updated_at FROM transactions WHERE created_at BETWEEN").
		WithArgs(startDate, endDate, pageSize, (page-1)*pageSize).
		WillReturnError(errors.New("database error"))

	// Call the function under test
	transactions, err := repo.RetrieveAllTransactions(page, pageSize, startDate, endDate)

	// Check for errors
	assert.Error(t, err)
	assert.Nil(t, transactions)

	// Verify that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCountAllTransactions_Success(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a new instance of the transactions repository
	repo := NewTransactionsRepository(db)

	// Mock the query and expected result
	startDate := "2024-03-01"
	endDate := "2024-03-10"
	expectedCount := 10

	// Expect the query and return the expected count
	mock.ExpectQuery("^SELECT COUNT(.+) FROM transactions WHERE created_at BETWEEN (.+) AND (.+)$").
		WithArgs(startDate, endDate).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedCount))

	// Call the function under test
	count, err := repo.CountAllTransactions(startDate, endDate)

	// Check for errors
	assert.NoError(t, err)

	// Check if the count matches the expected result
	assert.Equal(t, expectedCount, count)

	// Verify that all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCountAllTransactions_Failed(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a new instance of the transactions repository
	repo := NewTransactionsRepository(db)

	// Mocking the SQL query and simulating an error
	startDate := "2024-01-01"
	endDate := "2024-03-01"
	expectedError := errors.New("database error")
	mock.ExpectQuery("^SELECT COUNT(.+) FROM transactions WHERE created_at BETWEEN (.+) AND (.+)$").
		WithArgs(startDate, endDate).
		WillReturnError(expectedError)

	// Call the function under test
	count, err := repo.CountAllTransactions(startDate, endDate)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, 0, count)

	// Verify that the expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
