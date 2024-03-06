package Transactions

import "BookingRoom/model/dto/transactionDto"

type TransactionRepository interface {
	GetDailyTransactions(created_at string) ([]transactionDto.Transactions, error)
}

type TransactionUsecase interface {
	GetDailyTransaction(created_at string) ([]transactionDto.Transactions, error)
}
