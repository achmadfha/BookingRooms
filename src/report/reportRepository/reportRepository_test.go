package reportRepository

import (
	"BookingRoom/model/dto/reportDto"
	Report "BookingRoom/src/report"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ReportRepositoryTestSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo Report.ReportRepository
}

func (suite *ReportRepositoryTestSuite) SetupTest() {
	var err error
	suite.db, suite.mock, err = sqlmock.New()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	suite.repo = NewReportRepo(suite.db)
}

func (suite *ReportRepositoryTestSuite) TestGetAllReport_Success() {

	expectedReport := reportDto.Transactions{
		Transaction_id:  "11111111-1111-1111-1111-111111111111",
		FullName:        "Jane Smith",
		RoomName:        "Room 101",
		StartDate:       parseTime("2024-03-08"),
		EndDate:         parseTime("2024-03-08"),
		Description:     "Team building event",
		Created_at:      parseTime("2024-03-08"),
		Updated_at:      parseTime("2024-03-08"),
		Approved_by:     "John Doe",
		Approval_status: "DECLINE",
	}

	suite.mock.ExpectQuery(`SELECT 
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
			tl.updated_at;`).
		WithArgs("2024", "03", "06").
		WillReturnRows(sqlmock.NewRows([]string{"transaksi_id", "full_name", "room_name", "start_date", "end_date", "description", "created_at", "updated_at", "approved_by", "approval_status"}).
			AddRow("11111111-1111-1111-1111-111111111111", "Jane Smith", "Room 101", parseTime("2024-03-08"), parseTime("2024-03-08"),
				"Team building event", parseTime("2024-03-08"), parseTime("2024-03-08"), "John Doe", "DECLINE"))

	// Execute the method under test
	details, err := suite.repo.GetDailyTransactions("2024", "03", "06")

	// Verify the results
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedReport, details)

}

func parseTime(dateString string) time.Time {
	parsedTime, _ := time.Parse("2006-01-02", dateString)
	return parsedTime
}

func TestTransactionsRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ReportRepositoryTestSuite))
}
