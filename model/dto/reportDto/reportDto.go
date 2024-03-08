package reportDto

type Transactions struct {
	Transaction_id string `json:"transaksi_id"`
	Employee_id    string `json:"employee_id"`
	Room_id        string `json:"room_id"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	Created_at     string `json:"created_at"`
	Updated_at     string `json:"updated_at,omitempty"`
	Jumlah         string `json:"jumlah,omitempty"`
}
