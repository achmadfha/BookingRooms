package transactionDelivery

import (
	Transactions "BookingRoom/src/transactions"
	"encoding/csv"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionDelivery struct {
	transactionUC Transactions.TransactionUsecase
}

func NewTransactionDelivery(v1Group *gin.RouterGroup, transactionUC Transactions.TransactionUsecase) {
	handler := &TransactionDelivery{
		transactionUC: transactionUC,
	}

	transactionGroup := v1Group.Group("/transaction")
	{
		transactionGroup.GET("/report/daily/:created_at", handler.GetDaily)
		transactionGroup.GET("/report/daily/:created_at/export", handler.ExportDailyTransactionsCSV)

	}
}
func (h *TransactionDelivery) GetDaily(c *gin.Context) {
	created_at := c.Param("created_at")
	transactions, err := h.transactionUC.GetDailyTransaction(created_at)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
func (h *TransactionDelivery) ExportDailyTransactionsCSV(c *gin.Context) {
	directory := "./.csv/"
	created_at := c.Param("created_at")
	formatDate := time.Now().Format("2006-01-02")
	filename := directory + "daily_transactions_" + formatDate + ".csv"

	transactions, err := h.transactionUC.GetDailyTransaction(created_at)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Tulis header CSV
	if err := writer.Write([]string{"Transaction_id", "Employee_id", "Room_id", "StartDate", "Description", "EndDate", "Status", "Created_at", "Updated_at"}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tulis data transaksi
	for _, transaction := range transactions {
		if err := writer.Write([]string{transaction.Transaction_id, transaction.Employee_id,
			transaction.Room_id, transaction.StartDate, transaction.EndDate, transaction.Description,
			transaction.Status, transaction.Created_at, transaction.Updated_at}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=daily_transactions_"+created_at+".csv")
	http.ServeFile(c.Writer, c.Request, "daily_transactions_"+created_at+".csv")
}
