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
}

type TransactionsUseCase interface {
	RetrieveAllTransactions(page int, pageSize int, startDate string, endDate string) ([]transactionsDto.Transactions, json.Pagination, error)
	RetrieveTransactionsByID(trxID string) (transactionsDto.TransactionsDetailResponse, error)
	CreateTransactions(trxReq transactionsDto.TransactionsRequest) error
}
