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

	query := `SELECT
	  transaction_id,
	  employee_id,
	  room_id,
	  start_date,
	  end_date,
	  description,
	  status,
	  created_at,
	  updated_at
	FROM
	  transactions
	WHERE
	  created_at BETWEEN $1 AND $2
	LIMIT
	  $3 OFFSET $4`

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
			errors.New(fmt.Sprintf("error scanning transactions row: %s", err))
			continue
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("error iterating through transactions rows: %s", err))
	}

	return transactions, nil
}

func (t transactionsRepository) CountAllTransactions(startDate string, endDate string) (int, error) {
	var count int

	query := `SELECT
	  COUNT(*)
	FROM
	  transactions
	WHERE
	  created_at BETWEEN $1 AND $2`

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

	query := `INSERT INTO
	  transactions (
		transaction_id,
		employee_id,
		room_id,
		start_date,
		end_date,
		description,
		status
	  )
	VALUES
	  ($1, $2, $3, $4, $5, $6, $7) RETURNING transaction_id`

	var transactionID string
	err = tx.QueryRow(query, trx.ID, trx.EmployeeId, trx.RoomId, trx.StartDate, trx.EndDate, trx.Description, trx.Status).Scan(&transactionID)
	if err != nil {
		return err
	}

	qry := `INSERT INTO
	  transaction_logs (
		transaction_log_id,
		transaction_id,
		approval_status
	  )
	VALUES
	  ($1, $2, $3)`

	_, err = tx.Exec(qry, trx.TransactionsLogsID, transactionID, trx.Status)
	if err != nil {
		return err
	}

	return nil
}

func (t transactionsRepository) RetrieveRoomByID(roomID string) (transactionsDto.RoomResponse, error) {
	query := `SELECT
	  room_id,
	  room_details_id,
	  name,
	  status
	FROM
	  room
	WHERE
	  room_id = $1`

	var room transactionsDto.RoomResponse
	err := t.db.QueryRow(query, roomID).Scan(&room.ID, &room.RoomDetails, &room.Name, &room.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			// no rows
			return transactionsDto.RoomResponse{}, errors.New("01")
		}
		return transactionsDto.RoomResponse{}, err
	}

	return room, nil
}

func (t transactionsRepository) UpdateTrxLog(trxLog transactionsDto.TransactionLog) error {
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

	trxLogQuery := `UPDATE
	  transaction_logs
	SET
	  approved_by = $1,
	  approval_status = $2,
	  description = $3,
	  updated_at = CURRENT_TIMESTAMP
	WHERE
	  transaction_log_id = $4`

	_, err = tx.Exec(trxLogQuery, trxLog.ApprovedBy, trxLog.ApprovalStatus, trxLog.Descriptions, trxLog.TransactionLogID)
	if err != nil {
		return err
	}

	trxQuery := `UPDATE
	  transactions
	SET
	  status = $1,
	  updated_at = CURRENT_TIMESTAMP
	WHERE
	  transaction_id = $2`

	_, err = tx.Exec(trxQuery, trxLog.ApprovalStatus, trxLog.TransactionsID)
	if err != nil {
		return err
	}

	roomQuery := `UPDATE
	  room
	SET
	  status = 'BOOKED'
	WHERE
	  room_id = $1`

	_, err = tx.Exec(roomQuery, trxLog.RoomsID)
	if err != nil {
		return err
	}

	return nil
}

func (t transactionsRepository) RetrieveTrxLogDetailsByID(trxLogID string) (transactionsDto.TransactionLogDetailResponse, error) {
	var trxLogDetail transactionsDto.TransactionLogDetailResponse

	query := `SELECT 
        tl.transaction_log_id,
        t.transaction_id,
        t.employee_id,
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
        t.created_at AS transaction_created_at,
        t.updated_at AS transaction_updated_at,
        tl.approved_by,
        e2.full_name AS approved_by_name,
        e2.division AS approved_by_department,
        e2.phone_number AS approved_by_phone_number,
        e2.position AS approved_by_position,
        tl.approval_status,
        tl.description AS log_description,
        tl.created_at AS log_created_at,
        tl.updated_at AS log_updated_at
    FROM 
        transaction_logs tl
    JOIN 
        transactions t ON tl.transaction_id = t.transaction_id
    JOIN 
        employee e ON t.employee_id = e.employee_id
    JOIN 
        room r ON t.room_id = r.room_id
    JOIN 
        room_details rd ON r.room_details_id = rd.room_details_id
    JOIN 
        employee e2 ON tl.approved_by = e2.employee_id
    WHERE
        tl.transaction_log_id = $1`

	err := t.db.QueryRow(query, trxLogID).Scan(
		&trxLogDetail.TransactionLogID,
		&trxLogDetail.TransactionsID.ID,
		&trxLogDetail.TransactionsID.EmployeeId.ID,
		&trxLogDetail.TransactionsID.EmployeeId.FullName,
		&trxLogDetail.TransactionsID.EmployeeId.Division,
		&trxLogDetail.TransactionsID.EmployeeId.PhoneNumber,
		&trxLogDetail.TransactionsID.EmployeeId.Position,
		&trxLogDetail.TransactionsID.RoomId.ID,
		&trxLogDetail.TransactionsID.RoomId.Name,
		&trxLogDetail.TransactionsID.RoomId.RoomDetails.ID,
		&trxLogDetail.TransactionsID.RoomId.RoomDetails.RoomType,
		&trxLogDetail.TransactionsID.RoomId.RoomDetails.Capacity,
		pq.Array(&trxLogDetail.TransactionsID.RoomId.RoomDetails.Facility),
		&trxLogDetail.TransactionsID.RoomId.Status,
		&trxLogDetail.TransactionsID.StartDate,
		&trxLogDetail.TransactionsID.EndDate,
		&trxLogDetail.TransactionsID.Description,
		&trxLogDetail.TransactionsID.Status,
		&trxLogDetail.TransactionsID.CreatedAt,
		&trxLogDetail.TransactionsID.UpdatedAt,
		&trxLogDetail.ApprovedBy.ID,
		&trxLogDetail.ApprovedBy.FullName,
		&trxLogDetail.ApprovedBy.Division,
		&trxLogDetail.ApprovedBy.PhoneNumber,
		&trxLogDetail.ApprovedBy.Position,
		&trxLogDetail.ApprovalStatus,
		&trxLogDetail.Descriptions,
		&trxLogDetail.CreatedAt,
		&trxLogDetail.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// no rows
			return transactionsDto.TransactionLogDetailResponse{}, errors.New("01")
		}
		return transactionsDto.TransactionLogDetailResponse{}, err
	}

	return trxLogDetail, nil
}

func (t transactionsRepository) RetrieveTrxLogByID(trxLogID string) (transactionsDto.TransactionLogs, error) {
	query := `SELECT
	  transaction_log_id,
	  transaction_id
	FROM
	  transaction_logs
	WHERE
	  transaction_log_id = $1`

	var trxLog transactionsDto.TransactionLogs
	err := t.db.QueryRow(query, trxLogID).Scan(&trxLog.TransactionLogID, &trxLog.TransactionsID)
	if err != nil {
		if err == sql.ErrNoRows {
			// no rows
			return transactionsDto.TransactionLogs{}, errors.New("01")
		}
		return transactionsDto.TransactionLogs{}, err
	}

	return trxLog, err
}

func (t transactionsRepository) RetrieveAllTrxLog(page int, pageSize int, startDate string, endDate string) ([]transactionsDto.TransactionLogResponse, error) {
	offset := (page - 1) * pageSize
	limit := pageSize

	query := `SELECT
	  tl.transaction_log_id,
	  tl.transaction_id,
	  COALESCE(
		e.full_name,
		'pending for approval'
	  ) AS approved_by,
	  tl.approval_status,
	  COALESCE(
		tl.description,
		'pending for approval'
	  ) AS description,
	  tl.created_at,
	  tl.updated_at
	FROM
	  transaction_logs tl
	  LEFT JOIN employee e ON tl.approved_by = e.employee_id
	WHERE
	  tl.created_at BETWEEN $1 AND $2
	LIMIT
	  $3 OFFSET $4`

	rows, err := t.db.Query(query, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allTrxLogs []transactionsDto.TransactionLogResponse
	for rows.Next() {
		var allTrxLog transactionsDto.TransactionLogResponse

		err := rows.Scan(&allTrxLog.TransactionLogID, &allTrxLog.TransactionsID, &allTrxLog.ApprovedBy, &allTrxLog.ApprovalStatus, &allTrxLog.Descriptions, &allTrxLog.CreatedAt, &allTrxLog.UpdatedAt)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error scanning transactions row: %s", err))
		}

		allTrxLogs = append(allTrxLogs, allTrxLog)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("Error iterating through allTrxLogs rows: %s", err))
	}

	return allTrxLogs, nil
}

func (t transactionsRepository) CountAllTrxLogs(startDate string, endDate string) (int, error) {
	var count int

	query := `SELECT
	  COUNT(*)
	FROM
	  transaction_logs
	WHERE
	  created_at BETWEEN $1 AND $2`

	rows := t.db.QueryRow(query, startDate, endDate)
	err := rows.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
