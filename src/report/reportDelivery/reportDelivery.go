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
		reportGroup.GET("/daily/:year/:month/:day", handler.GetDaily)
		reportGroup.GET("/daily/:year/:month/:day/export", handler.ExportDailyTransactionsCSV)
		reportGroup.GET("/monthly/:year/:month", handler.GetMonthly)
		reportGroup.GET("/monthly/:year/:month/export", handler.ExportMonthlyTransactionsCSV)
		reportGroup.GET("/year/:year", handler.GetYear)
		reportGroup.GET("/year/:year/export", handler.ExportYearTransactionsCSV)

	}
}

func (h *ReportDelivery) GetDaily(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	day := c.Param("day")
	transactions, err := h.reportUC.GetDailyTransaction(year, month, day)
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
	year := c.Param("year")
	month := c.Param("month")
	day := c.Param("day")
	formatDate := time.Now().Format("2006-01-02")
	transactions, err := h.reportUC.GetDailyTransactionReport(year, month, day)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file := xlsx.NewFile()

	sheetName := "Daily Transactions " + formatDate
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	headers := []string{"Transaction_id", "Employee_id", "Room_id", "StartDate", "Description", "EndDate", "Status", "Created_at", "Updated_at", "Jumlah_Transaksi"}
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

	json.NewResponseSuccess(c, transactions, nil, "Data Berhasil Di Export", "01", "02")

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename="+excelFilename)
	http.ServeFile(c.Writer, c.Request, excelFilePath)
	c.Writer.Write(fileContent)

}

func (h *ReportDelivery) ExportMonthlyTransactionsCSV(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	formatDate := time.Now().Format("2006-01-02")
	transactions, err := h.reportUC.GetMonthlyTransactionReport(year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file := xlsx.NewFile()

	sheetName := "Monthly Transactions " + formatDate
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	headers := []string{"Transaction_id", "Employee_id", "Room_id", "StartDate", "Description", "EndDate", "Status", "Created_at", "Updated_at", "Jumlah_Transaksi"}
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
	excelFilename := "Monthly_transactions_" + formatDate + ".xlsx"
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

	json.NewResponseSuccess(c, nil, nil, "Data Berhasil Di Export", "02", "02")

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename="+excelFilename)
	http.ServeFile(c.Writer, c.Request, excelFilePath)
	c.Writer.Write(fileContent)

}

func (h *ReportDelivery) ExportYearTransactionsCSV(c *gin.Context) {
	year := c.Param("year")
	formatDate := time.Now().Format("2006-01-02")
	transactions, err := h.reportUC.GetYearTransactionReport(year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file := xlsx.NewFile()

	sheetName := "Year Transactions " + formatDate
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	headers := []string{"Transaction_id", "Employee_id", "Room_id", "StartDate", "Description", "EndDate", "Status", "Created_at", "Updated_at", "Jumlah_Transaksi"}
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
	excelFilename := "Year_transactions_" + formatDate + ".xlsx"
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

	json.NewResponseSuccess(c, nil, nil, "Data Berhasil Di Export", "02", "02")

	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename="+excelFilename)
	http.ServeFile(c.Writer, c.Request, excelFilePath)
	c.Writer.Write(fileContent)

}
