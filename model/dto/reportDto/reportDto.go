package reportDto

import "time"

type Transactions struct {
	Transaction_id string    `json:"transaksi_id"`
	Employee_id    string    `json:"employee_id"`
	Room_id        string    `json:"room_id"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at,omitempty"`
	Jumlah         string    `json:"jumlah,omitempty"`
}
