package transactionUsecase

import (
	"BookingRoom/model/dto/transactionDto"
	Transactions "BookingRoom/src/transactions"
)

type transactionUC struct {
	transactionRepo Transactions.TransactionRepository
}

func NewTransactionUsecase(transactionRepo Transactions.TransactionRepository) Transactions.TransactionUsecase {
	return &transactionUC{transactionRepo}
}

func (t *transactionUC) GetDailyTransaction(created_at string) ([]transactionDto.Transactions, error) {
	return t.transactionRepo.GetDailyTransactions(created_at)
}
