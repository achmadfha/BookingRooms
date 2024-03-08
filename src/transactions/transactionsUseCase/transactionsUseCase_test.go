package transactionsUseCase

//
//import (
//	"BookingRoom/model/dto/json"
//	"BookingRoom/model/dto/transactionsDto"
//	"errors"
//	"github.com/google/uuid"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"testing"
//	"time"
//)
//
//type MockTransactionsRepository struct {
//	mock.Mock
//}
//
//func (m *MockTransactionsRepository) RetrieveAllTransactions(page int, pageSize int, startDate string, endDate string) ([]transactionsDto.Transactions, error) {
//	// Mock the method and return some dummy data
//	args := m.Called(page, pageSize, startDate, endDate)
//	return args.Get(0).([]transactionsDto.Transactions), args.Error(1)
//}
//
//func (m *MockTransactionsRepository) CountAllTransactions(startDate string, endDate string) (int, error) {
//	// Mock the method and return some dummy data
//	args := m.Called(startDate, endDate)
//	return args.Int(0), args.Error(1)
//}
//
//func TestRetrieveAllTransactions_Success(t *testing.T) {
//	// Create an instance of the mock repository
//	mockRepo := new(MockTransactionsRepository)
//
//	// Mock data
//	createdAt := time.Now()
//	updatedAt := time.Now()
//
//	startDate := "2023-01-01"
//	endDate := "2023-12-31"
//	page := 1
//	pageSize := 1
//	transactionsData := []transactionsDto.Transactions{
//		{
//			ID:          uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
//			EmployeeId:  uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c9"),
//			RoomId:      uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c0"),
//			StartDate:   "2024-03-01",
//			EndDate:     "2024-03-05",
//			Description: "Test Transaction 1",
//			Status:      "completed",
//			CreatedAt:   createdAt,
//			UpdatedAt:   updatedAt,
//		},
//		{
//			ID:          uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"),
//			EmployeeId:  uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c9"),
//			RoomId:      uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c0"),
//			StartDate:   "2024-03-02",
//			EndDate:     "2024-03-06",
//			Description: "Test Transaction 2",
//			Status:      "pending",
//			CreatedAt:   createdAt,
//			UpdatedAt:   updatedAt,
//		},
//	}
//
//	mockRepo.On("RetrieveAllTransactions", page, pageSize, startDate, endDate).Return(transactionsData, nil)
//	mockRepo.On("CountAllTransactions", startDate, endDate).Return(2, nil)
//
//	useCase := NewTransactionsUseCase(mockRepo)
//
//	transactionsData, pagination, err := useCase.RetrieveAllTransactions(page, pageSize, startDate, endDate)
//
//	// Assert the results
//	assert.NoError(t, err)
//	assert.NotNil(t, transactionsData)
//	assert.Equal(t, len(transactionsData), len(transactionsData))
//	assert.Equal(t, json.Pagination{CurrentPage: 1, TotalPages: 2, TotalRecords: 2}, pagination)
//
//	// Assert that the mock's expectations were met
//	mockRepo.AssertExpectations(t)
//
//}
//
//func TestRetrieveAllTransactions_Failed(t *testing.T) {
//	// Create an instance of the mock repository
//	mockRepo := new(MockTransactionsRepository)
//
//	// Mock data
//	startDate := "2023-01-01"
//	endDate := "2023-12-31"
//	page := 1
//	pageSize := 1
//
//	// Mocking the repository method to return an error
//	mockRepo.On("RetrieveAllTransactions", page, pageSize, startDate, endDate).Return([]transactionsDto.Transactions{}, errors.New("error retrieving transactions"))
//
//	useCase := NewTransactionsUseCase(mockRepo)
//
//	transactionsData, pagination, err := useCase.RetrieveAllTransactions(page, pageSize, startDate, endDate)
//
//	// Assert the error
//	assert.Error(t, err)
//	assert.Nil(t, transactionsData)
//	assert.Equal(t, json.Pagination{}, pagination)
//
//	// Assert that the mock's expectations were met
//	mockRepo.AssertExpectations(t)
//}
