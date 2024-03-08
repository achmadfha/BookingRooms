package reportDelivery

import (
	"BookingRoom/model/dto/json"
	Report "BookingRoom/src/report"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type ReportDelivery struct {
	reportUC Report.ReportUsecase
}

func NewReportDelivery(v1Group *gin.RouterGroup, reportUC Report.ReportUsecase) {
	handler := &ReportDelivery{
		reportUC: reportUC,
	}

	reportGroup := v1Group.Group("/transaction/report")
	{
		reportGroup.GET("/daily/:created_at", handler.GetDaily)
		reportGroup.GET("/daily/:created_at/export", handler.ExportDailyTransactionsCSV)
		reportGroup.GET("/monthly/:year/:month", handler.GetMonthly)
		reportGroup.GET("/year/:year", handler.GetYear)

	}
}

func (h *ReportDelivery) GetDaily(c *gin.Context) {
	created_at := c.Param("created_at")
	transactions, err := h.reportUC.GetDailyTransaction(created_at)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	json.NewResponseSuccess(c, transactions, nil, "success", "01", "01")
}

func (h *ReportDelivery) GetMonthly(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	transactions, err := h.reportUC.GetMonthlyTransaction(year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	json.NewResponseSuccess(c, transactions, nil, "success", "02", "01")
}

func (h *ReportDelivery) GetYear(c *gin.Context) {
	year := c.Param("year")
	transactions, err := h.reportUC.GetYearTransaction(year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	json.NewResponseSuccess(c, transactions, nil, "success", "03", "01")
}

func (h *ReportDelivery) ExportDailyTransactionsCSV(c *gin.Context) {
	created_at := c.Param("created_at")
	formatDate := time.Now().Format("2006-01-02")
	transactions, err := h.reportUC.GetDailyTransaction(created_at)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Membuat file Excel baru
	file := xlsx.NewFile()

	// Membuat sheet baru
	sheetName := "Daily Transactions " + formatDate
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	headers := []string{"Transaction_id", "Employee_id", "Room_id", "StartDate", "Description", "EndDate", "Status", "Created_at", "Updated_at", "Jumlah"}
	headerRow := sheet.AddRow()
	for _, header := range headers {
		cell := headerRow.AddCell()
		cell.SetString(header)
	}

	for _, transaction := range transactions {
		dataRow := sheet.AddRow()
		dataRow.AddCell().SetString(transaction.Transaction_id)
		dataRow.AddCell().SetString(transaction.Employee_id)
		dataRow.AddCell().SetString(transaction.Room_id)
		dataRow.AddCell().SetString(transaction.StartDate)
		dataRow.AddCell().SetString(transaction.Description)
		dataRow.AddCell().SetString(transaction.EndDate)
		dataRow.AddCell().SetString(transaction.Status)
		dataRow.AddCell().SetString(transaction.Created_at)
		dataRow.AddCell().SetString(transaction.Updated_at)
		dataRow.AddCell().SetString(transaction.Jumlah)
	}

	excelDirectory := "./.excel"
	excelFilename := "daily_transactions_" + formatDate + ".xlsx"
	excelFilePath := filepath.Join(excelDirectory, excelFilename)
	if err := file.Save(excelFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileContent, err := os.ReadFile(excelFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename="+excelFilename)
	http.ServeFile(c.Writer, c.Request, excelFilePath)
	c.Writer.Write(fileContent)

}
