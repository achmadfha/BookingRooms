package reportRepository

import (
	"BookingRoom/model/dto/reportDto"
	Report "BookingRoom/src/report"
	"database/sql"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepo(db *sql.DB) Report.ReportRepository {
	return &ReportRepository{db}
}

func (t *ReportRepository) GetMostFrequentRoomNameDay(year, month, day string) (string, error) {
	var mostFrequentRoomName string

	query := `
		SELECT r.name
		FROM transaction_logs tl
		JOIN transactions t ON tl.transaction_id = t.transaction_id
		JOIN room r ON t.room_id = r.room_id
		WHERE EXTRACT(YEAR FROM tl.created_at) = $1 
			AND EXTRACT(MONTH FROM tl.created_at) = $2
			AND EXTRACT(DAY FROM tl.created_at) = $3
			AND tl.approval_status IN ('ACCEPT', 'DECLINE')
		GROUP BY r.name
		ORDER BY COUNT(*) DESC
		LIMIT 1
	`

	err := t.db.QueryRow(query, year, month, day).Scan(&mostFrequentRoomName)
	if err != nil {
		return "", err
	}

	return mostFrequentRoomName, nil
}

func (t *ReportRepository) GetDailyTransactionsReport(year, month, day string) ([]reportDto.Transactions, error) {
	query :=
		`SELECT 
	t.transaction_id AS transaksi_id,
    e_transaction.full_name AS fullname,  
	r.name AS room_name,
    t.start_date AS startdate,    
	t.description AS description,
    t.end_date AS enddate,    
	tl.approval_status AS approval_status_di_logs,
    e_logs.full_name AS approved_by,    
	tl.created_at AS created_at,
    tl.updated_at AS updated_at   
	FROM    transaction_logs tl
	JOIN    employee e_logs ON tl.approved_by = e_logs.employee_id
	JOIN    transactions t ON tl.transaction_id = t.transaction_id
	JOIN    employee e_transaction ON t.employee_id = e_transaction.employee_id
	JOIN    room r ON t.room_id = r.room_id
	WHERE 	EXTRACT(YEAR FROM tl.created_at) = $1 AND 
	EXTRACT(MONTH FROM tl.created_at) = $2 AND  
	EXTRACT(DAY FROM tl.created_at) = $3 AND
	tl.approval_status IN ('ACCEPT', 'DECLINE') 
	GROUP BY    
	t.transaction_id,
	e_transaction.full_name,   
	r.name,
	t.start_date,    
	t.description,
	t.end_date,    
	tl.approval_status,
	e_logs.full_name,    
	tl.created_at,
	tl.updated_at;`

	rows, err := t.db.Query(query, year, month, day)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.FullName, &transaction.RoomName,
			&transaction.StartDate, &transaction.Description, &transaction.EndDate, &transaction.Approval_status, &transaction.Approved_by,
			&transaction.Created_at, &transaction.Updated_at); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *ReportRepository) GetDailyTransactions(year, month, day string) ([]reportDto.Transactions, error) {
	query := `
	SELECT 
	t.transaction_id AS transaksi_id,
    e_transaction.full_name AS fullname,  
	r.name AS room_name,
    t.start_date AS startdate,    
	t.description AS description,
    t.end_date AS enddate,    
	tl.approval_status AS approval_status_di_logs,
    e_logs.full_name AS approved_by,    
	tl.created_at AS created_at,
    tl.updated_at AS updated_at   
	FROM    transaction_logs tl
	JOIN    employee e_logs ON tl.approved_by = e_logs.employee_id
	JOIN    transactions t ON tl.transaction_id = t.transaction_id
	JOIN    employee e_transaction ON t.employee_id = e_transaction.employee_id
	JOIN    room r ON t.room_id = r.room_id
	WHERE 	EXTRACT(YEAR FROM tl.created_at) = $1 AND 
	EXTRACT(MONTH FROM tl.created_at) = $2 AND  
	EXTRACT(DAY FROM tl.created_at) = $3 AND
	tl.approval_status IN ('ACCEPT', 'DECLINE') 
	GROUP BY    
			t.transaction_id,
			e_transaction.full_name,   
			r.name,
			t.start_date,    
			t.description,
			t.end_date,    
			tl.approval_status,
			e_logs.full_name,    
			tl.created_at,
			tl.updated_at;`

	rows, err := t.db.Query(query, year, month, day)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.FullName, &transaction.RoomName,
			&transaction.StartDate, &transaction.Description, &transaction.EndDate, &transaction.Approval_status, &transaction.Approved_by,
			&transaction.Created_at, &transaction.Updated_at); err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
func (t *ReportRepository) GetMostFrequentRoomNameMonth(year, month string) (string, error) {
	var mostFrequentRoomName string

	query := `
			SELECT r.name
			FROM transaction_logs tl
			JOIN transactions t ON tl.transaction_id = t.transaction_id
			JOIN room r ON t.room_id = r.room_id
			WHERE EXTRACT(YEAR FROM tl.created_at) = $1 
				AND EXTRACT(MONTH FROM tl.created_at) = $2
				AND tl.approval_status IN ('ACCEPT', 'DECLINE')
			GROUP BY r.name
			ORDER BY COUNT(*) DESC
			LIMIT 1
		`

	err := t.db.QueryRow(query, year, month).Scan(&mostFrequentRoomName)
	if err != nil {
		return "", err
	}

	return mostFrequentRoomName, nil
}

func (t *ReportRepository) GetMonthlyTransactionsReport(year, month string) ([]reportDto.Transactions, error) {
	query := `
	SELECT 
	t.transaction_id AS transaksi_id,
    e_transaction.full_name AS fullname,  
	r.name AS room_name,
    t.start_date AS startdate,    
	t.description AS description,
    t.end_date AS enddate,    
	tl.approval_status AS approval_status_di_logs,
    e_logs.full_name AS approved_by,    
	tl.created_at AS created_at,
    tl.updated_at AS updated_at   
	FROM    transaction_logs tl
	JOIN    employee e_logs ON tl.approved_by = e_logs.employee_id
	JOIN    transactions t ON tl.transaction_id = t.transaction_id
	JOIN    employee e_transaction ON t.employee_id = e_transaction.employee_id
	JOIN    room r ON t.room_id = r.room_id
	WHERE 	EXTRACT(YEAR FROM tl.created_at) = $1 AND 
	EXTRACT(MONTH FROM tl.created_at) = $2 AND  
	tl.approval_status IN ('ACCEPT', 'DECLINE') 
	GROUP BY    
	t.transaction_id,
	e_transaction.full_name,   
	r.name,
	t.start_date,    
	t.description,
	t.end_date,    
	tl.approval_status,
	e_logs.full_name,    
	tl.created_at,
	tl.updated_at;`

	rows, err := t.db.Query(query, year, month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.FullName, &transaction.RoomName,
			&transaction.StartDate, &transaction.Description, &transaction.EndDate, &transaction.Approval_status,
			&transaction.Approved_by, &transaction.Created_at, &transaction.Updated_at); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *ReportRepository) GetMonthlyTransactions(year, month string) ([]reportDto.Transactions, error) {
	query := `SELECT 
		t.transaction_id AS transaksi_id,
    	e_transaction.full_name AS fullname,  
		r.name AS room_name,
    	t.start_date AS startdate,    
		t.description AS description,
    	t.end_date AS enddate,    
		tl.approval_status AS approval_status_di_logs,
    	e_logs.full_name AS approved_by,    
		tl.created_at AS created_at,
    	tl.updated_at AS updated_at   
	FROM    transaction_logs tl
	JOIN    employee e_logs ON tl.approved_by = e_logs.employee_id
	JOIN    transactions t ON tl.transaction_id = t.transaction_id
	JOIN    employee e_transaction ON t.employee_id = e_transaction.employee_id
	JOIN    room r ON t.room_id = r.room_id
	WHERE 	EXTRACT(YEAR FROM tl.created_at) = $1 AND 
	 		EXTRACT(MONTH FROM tl.created_at) = $2 AND
	 		tl.approval_status IN ('ACCEPT', 'DECLINE') 
	GROUP BY    
			t.transaction_id,
			e_transaction.full_name,   
			r.name,
			t.start_date,    
			t.description,
			t.end_date,    
			tl.approval_status,
			e_logs.full_name,    
			tl.created_at,
			tl.updated_at`

	rows, err := t.db.Query(query, year, month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.FullName,
			&transaction.RoomName, &transaction.StartDate, &transaction.Description,
			&transaction.EndDate, &transaction.Approval_status, &transaction.Approved_by, &transaction.Created_at, &transaction.Updated_at); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}
func (t *ReportRepository) GetMostFrequentRoomNameYear(year string) (string, error) {
	var mostFrequentRoomName string

	query := `
			SELECT r.name
			FROM transaction_logs tl
			JOIN transactions t ON tl.transaction_id = t.transaction_id
			JOIN room r ON t.room_id = r.room_id
			WHERE EXTRACT(YEAR FROM tl.created_at) = $1 
				AND tl.approval_status IN ('ACCEPT', 'DECLINE')
			GROUP BY r.name
			ORDER BY COUNT(*) DESC
			LIMIT 1
		`

	err := t.db.QueryRow(query, year).Scan(&mostFrequentRoomName)
	if err != nil {
		return "", err
	}

	return mostFrequentRoomName, nil
}
func (t *ReportRepository) GetYearTransactionsReport(year string) ([]reportDto.Transactions, error) {
	query := `
	SELECT 
	t.transaction_id AS transaksi_id,
    e_transaction.full_name AS fullname,  
	r.name AS room_name,
    t.start_date AS startdate,    
	t.description AS description,
    t.end_date AS enddate,    
	tl.approval_status AS approval_status_di_logs,
    e_logs.full_name AS approved_by,    
	tl.created_at AS created_at,
    tl.updated_at AS updated_at   
	FROM    transaction_logs tl
	JOIN    employee e_logs ON tl.approved_by = e_logs.employee_id
	JOIN    transactions t ON tl.transaction_id = t.transaction_id
	JOIN    employee e_transaction ON t.employee_id = e_transaction.employee_id
	JOIN    room r ON t.room_id = r.room_id
	WHERE 	EXTRACT(YEAR FROM tl.created_at) = $1 AND 
	tl.approval_status IN ('ACCEPT', 'DECLINE') 
	GROUP BY    
	t.transaction_id,
	e_transaction.full_name,   
	r.name,
	t.start_date,    
	t.description,
	t.end_date,    
	tl.approval_status,
	e_logs.full_name,    
	tl.created_at,
	tl.updated_at;`

	rows, err := t.db.Query(query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.FullName, &transaction.RoomName,
			&transaction.StartDate, &transaction.Description, &transaction.EndDate, &transaction.Approval_status,
			&transaction.Approved_by, &transaction.Created_at, &transaction.Updated_at); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *ReportRepository) GetYearTransactions(year string) ([]reportDto.Transactions, error) {
	query := `SELECT 
	t.transaction_id AS transaksi_id,
	e_transaction.full_name AS fullname,  
	r.name AS room_name,
	t.start_date AS startdate,    
	t.description AS description,
	t.end_date AS enddate,    
	tl.approval_status AS approval_status_di_logs,
	e_logs.full_name AS approved_by,    
	tl.created_at AS created_at,
	tl.updated_at AS updated_at   
	FROM    transaction_logs tl
	JOIN    employee e_logs ON tl.approved_by = e_logs.employee_id
	JOIN    transactions t ON tl.transaction_id = t.transaction_id
	JOIN    employee e_transaction ON t.employee_id = e_transaction.employee_id
	JOIN    room r ON t.room_id = r.room_id
	WHERE 	EXTRACT(YEAR FROM tl.created_at) = $1 AND 
		 tl.approval_status IN ('ACCEPT', 'DECLINE') 
	GROUP BY    
		t.transaction_id,
		e_transaction.full_name,   
		r.name,
		t.start_date,    
		t.description,
		t.end_date,    
		tl.approval_status,
		e_logs.full_name,    
		tl.created_at,
		tl.updated_at`

	rows, err := t.db.Query(query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.FullName,
			&transaction.RoomName, &transaction.StartDate, &transaction.Description,
			&transaction.EndDate, &transaction.Approval_status, &transaction.Approved_by, &transaction.Created_at, &transaction.Updated_at); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}
