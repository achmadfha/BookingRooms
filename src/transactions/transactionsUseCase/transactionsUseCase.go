package transactionsUseCase

import (
	"BookingRoom/model/dto/json"
	"BookingRoom/model/dto/transactionsDto"
	"BookingRoom/src/employees"
	"BookingRoom/src/transactions"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math"
	"time"
)

type transactionsUC struct {
	transactionsRepository transactions.TransactionsRepository
	employeeRepository     employees.EmployeeRepository
}

func NewTransactionsUseCase(transactionsRepo transactions.TransactionsRepository, employeeRepo employees.EmployeeRepository) transactions.TransactionsUseCase {
	return &transactionsUC{transactionsRepo, employeeRepo}
}

func (t transactionsUC) RetrieveAllTransactions(page int, pageSize int, startDate string, endDate string) ([]transactionsDto.Transactions, json.Pagination, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 5
	}

	if startDate == "" || startDate == "0001-01-01" {
		// Set default value 1 years ago
		startDate = time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	}

	if endDate == "" || endDate == "0001-01-01" {
		// Set to time now
		endDate = time.Now().Format("2006-01-02")
	}

	transactionsData, err := t.transactionsRepository.RetrieveAllTransactions(page, pageSize, startDate, endDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, json.Pagination{}, errors.New("no rows found")
		}
		return nil, json.Pagination{}, err
	}

	for i := range transactionsData {
		startDate, err := time.Parse(time.RFC3339, transactionsData[i].StartDate)
		if err != nil {
			fmt.Println("Error parsing StartDate:", err)
			continue
		}
		endDate, err := time.Parse(time.RFC3339, transactionsData[i].EndDate)
		if err != nil {
			fmt.Println("Error parsing EndDate:", err)
			continue
		}

		createdAt, err := time.Parse(time.RFC3339Nano, transactionsData[i].CreatedAt)
		if err != nil {
			fmt.Println("Error parsing createdAt:", err)
			continue
		}

		updatedAt, err := time.Parse(time.RFC3339Nano, transactionsData[i].UpdatedAt)
		if err != nil {
			fmt.Println("Error parsing updatedAt:", err)
			continue
		}

		transactionsData[i].StartDate = startDate.Format("01-02-2006")
		transactionsData[i].EndDate = endDate.Format("01-02-2006")
		transactionsData[i].CreatedAt = createdAt.Format("02-01-2006 15:04:05")
		transactionsData[i].UpdatedAt = updatedAt.Format("02-01-2006 15:04:05")
	}

	totalTransactionsRows, err := t.transactionsRepository.CountAllTransactions(startDate, endDate)
	if err != nil {
		return nil, json.Pagination{}, err
	}

	totalPages := int(math.Ceil(float64(totalTransactionsRows) / float64(pageSize)))
	if page > totalPages {
		return nil, json.Pagination{}, err
	}

	if totalPages == 0 && totalTransactionsRows > 0 {
		totalPages = 1
	}

	pagination := json.Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: totalTransactionsRows,
	}

	return transactionsData, pagination, nil
}

func (t transactionsUC) RetrieveTransactionsByID(trxID string) (transactionsDto.TransactionsDetailResponse, error) {
	trxData, err := t.transactionsRepository.RetrieveTransactionsByID(trxID)
	if err != nil {
		if err.Error() == "01" {
			// 01 no rows
			return transactionsDto.TransactionsDetailResponse{}, errors.New("01")
		}
		return transactionsDto.TransactionsDetailResponse{}, err
	}

	startDate, err := time.Parse(time.RFC3339, trxData.StartDate)
	endDate, err := time.Parse(time.RFC3339, trxData.EndDate)
	createdAt, err := time.Parse(time.RFC3339Nano, trxData.CreatedAt)
	updatedAt, err := time.Parse(time.RFC3339Nano, trxData.UpdatedAt)

	trxData.StartDate = startDate.Format("01-02-2006")
	trxData.EndDate = endDate.Format("01-02-2006")
	trxData.CreatedAt = createdAt.Format("02-01-2006 15:04:05")
	trxData.UpdatedAt = updatedAt.Format("02-01-2006 15:04:05")

	return trxData, err
}

func (t transactionsUC) CreateTransactions(trxReq transactionsDto.TransactionsRequest) error {
	trxID, err := uuid.NewRandom()
	if err != nil {
		// error while generate uuid transaction
		return errors.New("01")
	}

	trxLogID, err := uuid.NewRandom()
	if err != nil {
		// error while generate uuid transaction logs
		return errors.New("02")
	}

	// todo
	// 1. check if employee and room id is existing
	employeID := trxReq.EmployeeId.String()
	employe, err := t.employeeRepository.RetrieveEmployeeById(employeID)
	if err != nil {
		if err.Error() == "02" {
			// employee not found
			return errors.New("03")
		}
		return err
	}

	roomID := trxReq.RoomId.String()
	rooms, err := t.transactionsRepository.RetrieveRoomByID(roomID)
	if err != nil {
		if err.Error() == "01" {
			// rooms not found
			return errors.New("04")
		}
		return err
	}

	// 2. checks if there room is available
	if rooms.Status != "AVAILABLE" {
		// room is not available
		return errors.New("05")
	}

	trxStatus := "PENDING"
	newTrx := transactionsDto.CreateTransactions{
		ID:                 trxID,
		EmployeeId:         employe.EmployeeId,
		RoomId:             rooms.ID,
		StartDate:          trxReq.StartDate,
		EndDate:            trxReq.EndDate,
		Description:        trxReq.Description,
		Status:             trxStatus,
		TransactionsLogsID: trxLogID,
	}

	err = t.transactionsRepository.CreateTransactions(newTrx)
	if err != nil {
		return err
	}

	return nil
}

func (t transactionsUC) UpdateTrxLog(trxLog transactionsDto.TransactionLog) error {
	// check employee id are exists
	employeID := trxLog.ApprovedBy.String()
	employeIDExists, err := t.employeeRepository.RetrieveEmployeeById(employeID)
	if err != nil {
		if err.Error() == "02" {
			// employee not found
			return errors.New("01")
		}
		return err
	}

	// check trx log id are exists
	trxLogID := trxLog.TransactionLogID.String()
	trxLogIDExists, err := t.transactionsRepository.RetrieveTrxLogByID(trxLogID)
	if err != nil {
		if err.Error() == "01" {
			// trx log id not found
			return errors.New("02")
		}
		return err
	}

	// do updated
	trxData := transactionsDto.TransactionLog{
		ApprovedBy:       employeIDExists.EmployeeId,
		ApprovalStatus:   trxLog.ApprovalStatus,
		Descriptions:     trxLog.Descriptions,
		TransactionLogID: trxLogIDExists.TransactionLogID,
		TransactionsID:   trxLogIDExists.TransactionsID,
	}

	err = t.transactionsRepository.UpdateTrxLog(trxData)
	if err != nil {
		return err
	}

	return nil
}

func (t transactionsUC) RetrieveTrxLogByID(trxLodID string) (transactionsDto.TransactionLogDetailResponse, error) {
	trxLogData, err := t.transactionsRepository.RetrieveTrxLogDetailsByID(trxLodID)
	if err != nil {
		if err.Error() == "01" {
			// 01 no rows
			return transactionsDto.TransactionLogDetailResponse{}, errors.New("01")
		}
		return transactionsDto.TransactionLogDetailResponse{}, err
	}

	startDate, err := time.Parse(time.RFC3339, trxLogData.TransactionsID.StartDate)
	endDate, err := time.Parse(time.RFC3339, trxLogData.TransactionsID.EndDate)
	createdAtTrxDetails, err := time.Parse(time.RFC3339Nano, trxLogData.TransactionsID.CreatedAt)
	updatedAtTrxDetails, err := time.Parse(time.RFC3339Nano, trxLogData.TransactionsID.UpdatedAt)
	createdAt, err := time.Parse(time.RFC3339Nano, trxLogData.CreatedAt)
	updatedAt, err := time.Parse(time.RFC3339Nano, trxLogData.UpdatedAt)

	trxLogData.TransactionsID.StartDate = startDate.Format("01-02-2006")
	trxLogData.TransactionsID.EndDate = endDate.Format("01-02-2006")
	trxLogData.TransactionsID.CreatedAt = createdAtTrxDetails.Format("02-01-2006 15:04:05")
	trxLogData.TransactionsID.UpdatedAt = updatedAtTrxDetails.Format("02-01-2006 15:04:05")
	trxLogData.CreatedAt = createdAt.Format("02-01-2006 15:04:05")
	trxLogData.UpdatedAt = updatedAt.Format("02-01-2006 15:04:05")

	return trxLogData, err
}

func (t transactionsUC) RetrieveAllTrxLog(page int, pageSize int, startDate string, endDate string) ([]transactionsDto.TransactionLog, json.Pagination, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 5
	}

	if startDate == "" || startDate == "0001-01-01" {
		// Set default value 1 years ago
		startDate = time.Now().AddDate(-1, 0, 0).Format("2006-01-02")
	}

	if endDate == "" || endDate == "0001-01-01" {
		// Set to time now
		endDate = time.Now().Format("2006-01-02")
	}

	transactionsLogData, err := t.transactionsRepository.RetrieveAllTrxLog(page, pageSize, startDate, endDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, json.Pagination{}, errors.New("no rows found")
		}
		return nil, json.Pagination{}, err
	}

	for i := range transactionsLogData {
		createdAt, err := time.Parse(time.RFC3339Nano, transactionsLogData[i].CreatedAt)
		if err != nil {
			fmt.Println("Error parsing createdAt:", err)
			continue
		}

		updatedAt, err := time.Parse(time.RFC3339Nano, transactionsLogData[i].UpdatedAt)
		if err != nil {
			fmt.Println("Error parsing updatedAt:", err)
			continue
		}

		transactionsLogData[i].CreatedAt = createdAt.Format("02-01-2006 15:04:05")
		transactionsLogData[i].UpdatedAt = updatedAt.Format("02-01-2006 15:04:05")
	}

	totalTransactionLogsRows, err := t.transactionsRepository.CountAllTrxLogs(startDate, endDate)
	if err != nil {
		return nil, json.Pagination{}, err
	}

	totalPages := int(math.Ceil(float64(totalTransactionLogsRows) / float64(pageSize)))
	if page > totalPages {
		return nil, json.Pagination{}, err
	}

	if totalPages == 0 && totalTransactionLogsRows > 0 {
		totalPages = 1
	}

	pagination := json.Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: totalTransactionLogsRows,
	}

	return transactionsLogData, pagination, nil
}
