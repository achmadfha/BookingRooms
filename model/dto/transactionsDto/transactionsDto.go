package transactionsDto

import (
	"github.com/google/uuid"
	"time"
)

type (
	Transactions struct {
		ID          uuid.UUID `json:"transaction_id"`
		EmployeeId  uuid.UUID `json:"employee_id"`
		RoomId      uuid.UUID `json:"room_id"`
		StartDate   string    `json:"start_date"`
		EndDate     string    `json:"end_date"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at,omitempty"`
	}

	TransactionsRequest struct {
		EmployeeId  uuid.UUID `json:"employee_id"`
		RoomId      uuid.UUID `json:"room_id"`
		StartDate   string    `json:"start_date"`
		EndDate     string    `json:"end_date"`
		Description string    `json:"description"`
	}

	CreateTransactions struct {
		ID                 uuid.UUID `json:"transaction_id"`
		EmployeeId         uuid.UUID `json:"employee_id"`
		RoomId             uuid.UUID `json:"room_id"`
		StartDate          string    `json:"start_date"`
		EndDate            string    `json:"end_date"`
		Description        string    `json:"description"`
		Status             string    `json:"status"`
		TransactionsLogsID uuid.UUID `json:"transactions_log_id"`
	}

	TransactionsDetailResponse struct {
		ID          uuid.UUID       `json:"transaction_id"`
		EmployeeId  EmployeeDetails `json:"employee"`
		RoomId      Rooms           `json:"room"`
		StartDate   string          `json:"start_date"`
		EndDate     string          `json:"end_date"`
		Description string          `json:"description"`
		Status      string          `json:"status"`
		CreatedAt   time.Time       `json:"created_at"`
		UpdatedAt   time.Time       `json:"updated_at"`
	}

	EmployeeDetails struct {
		ID          uuid.UUID `json:"employee_id"`
		FullName    string    `json:"full_name"`
		Division    string    `json:"division"`
		PhoneNumber string    `json:"phone_number"`
		Position    string    `json:"position"`
	}

	Rooms struct {
		ID          uuid.UUID   `json:"room_id"`
		RoomDetails RoomDetails `json:"room_details"`
		Name        string      `json:"name"`
		Status      string      `json:"status"`
	}

	RoomDetails struct {
		ID       uuid.UUID `json:"room_details_id"`
		RoomType string    `json:"room_type"`
		Capacity int       `json:"capacity"`
		Facility []string  `json:"facility"`
	}
)
