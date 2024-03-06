package transactionRepository

import (
	"BookingRoom/model/dto/transactionDto"
	Transactions "BookingRoom/src/transactions"
	"database/sql"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) Transactions.TransactionRepository {
	return &TransactionRepository{db}
}

func (t *TransactionRepository) GetDailyTransactions(created_at string) ([]transactionDto.Transactions, error) {
	query := `SELECT * from transactions WHERE created_at = $1`

	rows, err := t.db.Query(query, created_at)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []transactionDto.Transactions

	for rows.Next() {
		var transaction transactionDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.Employee_id, &transaction.Room_id,
			&transaction.StartDate, &transaction.Description, &transaction.EndDate, &transaction.Status, &transaction.Created_at, &transaction.Updated_at); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
