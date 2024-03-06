package transactionsRepository

import (
	"BookingRoom/model/dto/transactionsDto"
	"BookingRoom/src/transactions"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
)

type transactionsRepository struct {
	db *sql.DB
}

func NewTransactionsRepository(db *sql.DB) transactions.TransactionsRepository {
	return &transactionsRepository{db}
}

func (t transactionsRepository) RetrieveAllTransactions(page int, pageSize int, startDate string, endDate string) ([]transactionsDto.Transactions, error) {
	offset := (page - 1) * pageSize
	limit := pageSize

	query := `SELECT transaction_id, employee_id, room_id, start_date, end_date, description, status, created_at, updated_at FROM transactions WHERE created_at BETWEEN $1 and $2 LIMIT $3 OFFSET $4`

	rows, err := t.db.Query(query, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []transactionsDto.Transactions
	for rows.Next() {
		var transaction transactionsDto.Transactions
		err := rows.Scan(&transaction.ID, &transaction.EmployeeId, &transaction.RoomId, &transaction.StartDate, &transaction.EndDate, &transaction.Description, &transaction.Status, &transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			errors.New(fmt.Sprintf("Error scanning transactions row: %s", err))
			continue
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("Error iterating through transactions rows: %s", err))
	}

	return transactions, nil
}

func (t transactionsRepository) CountAllTransactions(startDate string, endDate string) (int, error) {
	var count int

	query := `SELECT COUNT(*) FROM transactions WHERE created_at BETWEEN $1 AND $2`

	rows := t.db.QueryRow(query, startDate, endDate)
	err := rows.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (t transactionsRepository) RetrieveTransactionsByID(trxID string) (transactionsDto.TransactionsDetailResponse, error) {
	var trxDetails transactionsDto.TransactionsDetailResponse

	query := `SELECT 
    t.transaction_id,
    e.employee_id,
    e.full_name AS employee_name,
    e.division AS department,
    e.phone_number,
    e.position,
    r.room_id,
    r.name AS room_name,
    rd.room_details_id,
    rd.room_type,
    rd.capacity,
    rd.facility,
    r.status AS room_status,
    t.start_date,
    t.end_date,
    t.description,
    t.status AS transaction_status,
    t.created_at,
    t.updated_at
FROM 
    transactions t
JOIN 
    employee e ON t.employee_id = e.employee_id
JOIN 
    room r ON t.room_id = r.room_id
JOIN 
    room_details rd ON r.room_details_id = rd.room_details_id
WHERE
    t.transaction_id = $1`

	err := t.db.QueryRow(query, trxID).Scan(
		&trxDetails.ID,
		&trxDetails.EmployeeId.ID,
		&trxDetails.EmployeeId.FullName,
		&trxDetails.EmployeeId.Division,
		&trxDetails.EmployeeId.PhoneNumber,
		&trxDetails.EmployeeId.Position,
		&trxDetails.RoomId.ID,
		&trxDetails.RoomId.Name,
		&trxDetails.RoomId.RoomDetails.ID,
		&trxDetails.RoomId.RoomDetails.RoomType,
		&trxDetails.RoomId.RoomDetails.Capacity,
		pq.Array(&trxDetails.RoomId.RoomDetails.Facility),
		&trxDetails.RoomId.Status,
		&trxDetails.StartDate,
		&trxDetails.EndDate,
		&trxDetails.Description,
		&trxDetails.Status,
		&trxDetails.CreatedAt,
		&trxDetails.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return transactionsDto.TransactionsDetailResponse{}, errors.New("01")
		}
		return transactionsDto.TransactionsDetailResponse{}, err
	}

	fmt.Println(trxDetails)
	return trxDetails, nil
}

func (t transactionsRepository) CreateTransactions(trx transactionsDto.CreateTransactions) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// insert transactions
	query := `INSERT INTO transactions (transaction_id, employee_id, room_id, start_date, end_date, description, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING transaction_id`

	var transactionID string
	err = tx.QueryRow(query, trx.ID, trx.EmployeeId, trx.RoomId, trx.StartDate, trx.EndDate, trx.Description, trx.Status).Scan(&transactionID)
	if err != nil {
		return err
	}

	// insert transactions log
	qry := `INSERT INTO transaction_logs (transaction_log_id, transaction_id, approval_status) VALUES ($1, $2, $3)`

	_, err = tx.Exec(qry, trx.TransactionsLogsID, transactionID, trx.Status)
	if err != nil {
		return err
	}

	return nil
}
