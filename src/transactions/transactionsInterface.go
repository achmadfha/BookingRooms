package transactions

import (
	"BookingRoom/model/dto/json"
	"BookingRoom/model/dto/transactionsDto"
)

type TransactionsRepository interface {
	RetrieveAllTransactions(page int, pageSize int, startDate string, endDate string) ([]transactionsDto.Transactions, error)
	CountAllTransactions(startDate string, endDate string) (int, error)
	RetrieveTransactionsByID(trxID string) (transactionsDto.TransactionsDetailResponse, error)
	CreateTransactions(trx transactionsDto.CreateTransactions) error
	RetrieveRoomByID(roomID string) (transactionsDto.RoomResponse, error)
	UpdateTrxLog(trxLog transactionsDto.TransactionLog) error
	RetrieveTrxLogDetailsByID(trxLogID string) (transactionsDto.TransactionLogDetailResponse, error)
	RetrieveTrxLogByID(trxLogID string) (transactionsDto.TransactionLogs, error)
	RetrieveAllTrxLog(page int, pageSize int, startDate string, endDate string) ([]transactionsDto.TransactionLog, error)
	CountAllTrxLogs(startDate string, endDate string) (int, error)
}

type TransactionsUseCase interface {
	RetrieveAllTransactions(page int, pageSize int, startDate string, endDate string) ([]transactionsDto.Transactions, json.Pagination, error)
	RetrieveTransactionsByID(trxID string) (transactionsDto.TransactionsDetailResponse, error)
	CreateTransactions(trxReq transactionsDto.TransactionsRequest) error
	UpdateTrxLog(trxLog transactionsDto.TransactionLog) error
	RetrieveTrxLogByID(trxLodID string) (transactionsDto.TransactionLogDetailResponse, error)
	RetrieveAllTrxLog(page int, pageSize int, startDate string, endDate string) ([]transactionsDto.TransactionLog, json.Pagination, error)
}
