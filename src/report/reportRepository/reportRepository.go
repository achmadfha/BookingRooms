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

func (t *ReportRepository) GetDailyTransactionsReport(year, month, day string) ([]reportDto.Transactions, error) {
	query := `
	SELECT 
		'00000000-0000-0000-0000-000000000000'::uuid AS transaction_id, 
		'00000000-0000-0000-0000-000000000000'::uuid AS employee_id, 
		'00000000-0000-0000-0000-000000000000'::uuid AS room_id, 
		'0001-01-01'::timestamp AS start_date, 
		'0001-01-01'::timestamp AS end_date, 
		'' AS description, 
		COALESCE(status, 'PENDING') AS status, 
		'0001-01-01'::timestamp AS created_at, 
		'0001-01-01'::timestamp AS updated_at,
		COUNT(created_at) AS Subtotal
	FROM transactions 
	WHERE EXTRACT(YEAR FROM created_at) = $1 AND EXTRACT(MONTH FROM created_at) = $2 
	AND EXTRACT(DAY FROM created_at) = $3  AND status IN ('ACCEPT', 'DECLINE')
	GROUP BY created_at, status
	
	UNION ALL
	
	SELECT 
		transaction_id, 
		employee_id, 
		room_id, 
		start_date, 
		end_date, 
		description, 
		status, 
		created_at, 
		updated_at,
		0 AS Subtotal
	FROM transactions
	WHERE EXTRACT(YEAR FROM created_at) = $1 AND EXTRACT(MONTH FROM created_at) = $2 
	AND EXTRACT(DAY FROM created_at) = $3  AND status IN ('ACCEPT', 'DECLINE');`

	rows, err := t.db.Query(query, year, month, day)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.Employee_id, &transaction.Room_id,
			&transaction.StartDate, &transaction.EndDate, &transaction.Description, &transaction.Status,
			&transaction.Created_at, &transaction.Updated_at, &transaction.Jumlah); err != nil {
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
		transaction_id, 
		employee_id, 
		room_id, 
		start_date, 
		end_date, 
		description, 
		status, 
		created_at, 
		updated_at
	FROM transactions
	WHERE EXTRACT(YEAR FROM created_at) = $1 AND EXTRACT(MONTH FROM created_at) = $2 
	AND EXTRACT(DAY FROM created_at) = $3;`

	rows, err := t.db.Query(query, year, month, day)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.Employee_id, &transaction.Room_id,
			&transaction.StartDate, &transaction.EndDate, &transaction.Description, &transaction.Status,
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

func (t *ReportRepository) GetMonthlyTransactionsReport(year, month string) ([]reportDto.Transactions, error) {
	query := `
	SELECT 
		'00000000-0000-0000-0000-000000000000'::uuid AS transaction_id, 
		'00000000-0000-0000-0000-000000000000'::uuid AS employee_id, 
		'00000000-0000-0000-0000-000000000000'::uuid AS room_id, 
		'0001-01-01'::timestamp AS start_date, 
		'0001-01-01'::timestamp AS end_date, 
		'' AS description, 
		COALESCE(status, 'PENDING') AS status, 
		'0001-01-01'::timestamp AS created_at, 
		'0001-01-01'::timestamp AS updated_at,
		COUNT(created_at) AS Subtotal
	FROM transactions 
	WHERE EXTRACT(YEAR FROM created_at) = $1 AND EXTRACT(MONTH FROM created_at) = $2 AND status IN ('ACCEPT', 'DECLINE')
	GROUP BY created_at, status
	
	UNION ALL
	
	SELECT 
		transaction_id, 
		employee_id, 
		room_id, 
		start_date, 
		end_date, 
		description, 
		status, 
		created_at, 
		updated_at,
		0 AS Subtotal
	FROM transactions
	WHERE EXTRACT(YEAR FROM created_at) = $1 AND EXTRACT(MONTH FROM created_at) = $2 AND status IN ('ACCEPT', 'DECLINE');`

	rows, err := t.db.Query(query, year, month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.Employee_id, &transaction.Room_id,
			&transaction.StartDate, &transaction.EndDate, &transaction.Description, &transaction.Status,
			&transaction.Created_at, &transaction.Updated_at, &transaction.Jumlah); err != nil {
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
	query := `SELECT * from transactions WHERE EXTRACT(YEAR FROM created_at) = $1 AND EXTRACT(MONTH FROM created_at) = $2`

	rows, err := t.db.Query(query, year, month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.Employee_id,
			&transaction.Room_id, &transaction.StartDate, &transaction.EndDate,
			&transaction.Description, &transaction.Status, &transaction.Created_at, &transaction.Updated_at); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *ReportRepository) GetYearTransactionsReport(year string) ([]reportDto.Transactions, error) {
	query := `
	SELECT 
		'00000000-0000-0000-0000-000000000000'::uuid AS transaction_id, 
		'00000000-0000-0000-0000-000000000000'::uuid AS employee_id, 
		'00000000-0000-0000-0000-000000000000'::uuid AS room_id, 
		'0001-01-01'::timestamp AS start_date, 
		'0001-01-01'::timestamp AS end_date, 
		'' AS description, 
		COALESCE(status, 'PENDING') AS status, 
		'0001-01-01'::timestamp AS created_at, 
		'0001-01-01'::timestamp AS updated_at,
		COUNT(created_at) AS Subtotal
	FROM transactions 
	WHERE EXTRACT(YEAR FROM created_at) = $1 AND status IN ('ACCEPT', 'DECLINE')
	GROUP BY created_at, status
	
	UNION ALL
	
	SELECT 
		transaction_id, 
		employee_id, 
		room_id, 
		start_date, 
		end_date, 
		description, 
		status, 
		created_at, 
		updated_at,
		0 AS Subtotal
	FROM transactions
	WHERE EXTRACT(YEAR FROM created_at) = $1 AND status IN ('ACCEPT', 'DECLINE');`

	rows, err := t.db.Query(query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.Employee_id, &transaction.Room_id,
			&transaction.StartDate, &transaction.EndDate, &transaction.Description, &transaction.Status,
			&transaction.Created_at, &transaction.Updated_at, &transaction.Jumlah); err != nil {
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
	query := `SELECT * from transactions WHERE EXTRACT(YEAR FROM created_at) = $1`

	rows, err := t.db.Query(query, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []reportDto.Transactions

	for rows.Next() {
		var transaction reportDto.Transactions
		if err := rows.Scan(&transaction.Transaction_id, &transaction.Employee_id,
			&transaction.Room_id, &transaction.StartDate, &transaction.EndDate,
			&transaction.Description, &transaction.Status, &transaction.Created_at, &transaction.Updated_at); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}
