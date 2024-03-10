package reportDto

import "time"

type Transactions struct {
	Transaction_id  string    `json:"transaksi_id,omitempty"`
	Employee_id     string    `json:"employee_id,omitempty"`
	FullName        string    `json:"full_name,omitempty"`
	Room_id         string    `json:"room_id,omitempty"`
	RoomName        string    `json:"room_name,omitempty"`
	StartDate       time.Time `json:"start_date,omitempty"`
	EndDate         time.Time `json:"end_date,omitempty"`
	Description     string    `json:"description,omitempty"`
	Status          string    `json:"status,omitempty,omitempty"`
	Created_at      time.Time `json:"created_at,omitempty"`
	Updated_at      time.Time `json:"updated_at,omitempty"`
	Approved_by     string    `json:"approved_by,omitempty"`
	Approval_status string    `json:"approval_status"`
}
