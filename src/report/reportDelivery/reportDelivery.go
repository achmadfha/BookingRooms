package reportDelivery

import (
	"BookingRoom/model/dto/json"
	"BookingRoom/pkg/middleware"
	Report "BookingRoom/src/report"
	"fmt"
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
		reportGroup.GET("/daily/:year/:month/:day", middleware.JWTAuth("ADMIN", "GA"), handler.GetDaily)
		reportGroup.GET("/daily/:year/:month/:day/export", middleware.JWTAuth("ADMIN", "GA"), handler.ExportDailyTransactionsCSV)
		reportGroup.GET("/monthly/:year/:month", middleware.JWTAuth("ADMIN", "GA"), handler.GetMonthly)
		reportGroup.GET("/monthly/:year/:month/export", middleware.JWTAuth("ADMIN", "GA"), handler.ExportMonthlyTransactionsCSV)
		reportGroup.GET("/year/:year", middleware.JWTAuth("ADMIN", "GA"), handler.GetYear)
		reportGroup.GET("/year/:year/export", middleware.JWTAuth("ADMIN", "GA"), handler.ExportYearTransactionsCSV)

	}
}

func (h *ReportDelivery) GetDaily(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	day := c.Param("day")
	transactions, err := h.reportUC.GetDailyTransaction(year, month, day)
	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(c, transactions, nil, "success", "01", "01")
}

func (h *ReportDelivery) GetMonthly(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	transactions, err := h.reportUC.GetMonthlyTransaction(year, month)
	if err != nil {
		json.NewResponseError(c, err.Error(), "02", "01")
		return
	}

	json.NewResponseSuccess(c, transactions, nil, "success", "02", "01")
}

func (h *ReportDelivery) GetYear(c *gin.Context) {
	year := c.Param("year")
	transactions, err := h.reportUC.GetYearTransaction(year)
	if err != nil {
		json.NewResponseError(c, err.Error(), "03", "01")
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
		json.NewResponseError(c, err.Error(), "01", "02")
		return
	}

	mostCommonRoomName, err := h.reportUC.GetMostFrequentRoomNamesDay(year, month, day)
	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "03")
		return
	}

	file := xlsx.NewFile()

	sheetName := "Daily Transactions " + formatDate
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "04")
		return
	}

	headers := []string{"Transaction_id", "FullName", "Room Name", "StartDate", "Description", "EndDate", "Approval Status", "Approved BY", "Created_at", "Updated_at"}
	headerRow := sheet.AddRow()
	for _, header := range headers {
		cell := headerRow.AddCell()
		cell.SetString(header)
	}

	totalTransactions := 0

	for _, transaction := range transactions {
		dataRow := sheet.AddRow()
		dataRow.AddCell().SetString(transaction.Transaction_id)
		dataRow.AddCell().SetString(transaction.Employee_id)
		dataRow.AddCell().SetString(transaction.Room_id)
		dataRow.AddCell().SetString(transaction.StartDate.Format("2006-01-02"))
		dataRow.AddCell().SetString(transaction.Description)
		dataRow.AddCell().SetString(transaction.EndDate.Format("2006-01-02"))
		dataRow.AddCell().SetString(transaction.Approval_status)
		dataRow.AddCell().SetString(transaction.Approved_by)
		dataRow.AddCell().SetString(transaction.Created_at.Format("2006-01-02"))
		dataRow.AddCell().SetString(transaction.Updated_at.Format("2006-01-02"))

		totalTransactions++

	}

	totalRow := sheet.AddRow()
	totalRow.AddCell().SetString("TOTAL TRANSACTION")
	totalRow.AddCell().SetString(fmt.Sprintf("%d", totalTransactions))

	mostCommonRoomRow := sheet.AddRow()
	mostCommonRoomRow.AddCell().SetString("COMMON ROOM ")
	mostCommonRoomRow.AddCell().SetString(mostCommonRoomName)

	excelDirectory := "./.excel"
	excelFilename := "daily_transactions_" + formatDate + ".xlsx"
	excelFilePath := filepath.Join(excelDirectory, excelFilename)
	if err := file.Save(excelFilePath); err != nil {
		json.NewResponseError(c, err.Error(), "01", "05")
		return
	}

	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "06")
		return
	}

	json.NewResponseSuccess(c, nil, nil, "Data Berhasil Di Export", "01", "02")

}

func (h *ReportDelivery) ExportMonthlyTransactionsCSV(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	formatDate := time.Now().Format("2006-01-02")
	transactions, err := h.reportUC.GetMonthlyTransactionReport(year, month)
	if err != nil {
		json.NewResponseError(c, err.Error(), "02", "02")
		return
	}

	mostCommonRoomName, err := h.reportUC.GetMostFrequentRoomNameMonths(year, month)
	if err != nil {
		json.NewResponseError(c, err.Error(), "02", "03")
		return
	}

	file := xlsx.NewFile()

	sheetName := "Monthly Transactions " + formatDate
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		json.NewResponseError(c, err.Error(), "02", "04")
		return
	}

	headers := []string{"Transaction_id", "FullName", "Room Name", "StartDate", "Description", "EndDate", "Approval Status", "Approved BY", "Created_at", "Updated_at"}
	headerRow := sheet.AddRow()
	for _, header := range headers {
		cell := headerRow.AddCell()
		cell.SetString(header)
	}

	totalTransactions := 0

	for _, transaction := range transactions {
		dataRow := sheet.AddRow()
		dataRow.AddCell().SetString(transaction.Transaction_id)
		dataRow.AddCell().SetString(transaction.Employee_id)
		dataRow.AddCell().SetString(transaction.Room_id)
		dataRow.AddCell().SetString(transaction.StartDate.Format("2006-01-02"))
		dataRow.AddCell().SetString(transaction.Description)
		dataRow.AddCell().SetString(transaction.EndDate.Format("2006-01-02"))
		dataRow.AddCell().SetString(transaction.Approval_status)
		dataRow.AddCell().SetString(transaction.Approved_by)
		dataRow.AddCell().SetString(transaction.Created_at.Format("2006-01-02"))
		dataRow.AddCell().SetString(transaction.Updated_at.Format("2006-01-02"))

		totalTransactions++
	}

	totalRow := sheet.AddRow()
	totalRow.AddCell().SetString("TOTAL TRANSACTION")
	totalRow.AddCell().SetString(fmt.Sprintf("%d", totalTransactions))

	mostCommonRoomRow := sheet.AddRow()
	mostCommonRoomRow.AddCell().SetString("COMMON ROOM")
	mostCommonRoomRow.AddCell().SetString(mostCommonRoomName)

	excelDirectory := "./.excel"
	excelFilename := "Monthly_transactions_" + formatDate + ".xlsx"
	excelFilePath := filepath.Join(excelDirectory, excelFilename)
	if err := file.Save(excelFilePath); err != nil {
		json.NewResponseError(c, err.Error(), "02", "05")
		return
	}

	if err != nil {
		json.NewResponseError(c, err.Error(), "02", "06")
		return
	}

	json.NewResponseSuccess(c, nil, nil, "Data Berhasil Di Export", "02", "02")

}

func (h *ReportDelivery) ExportYearTransactionsCSV(c *gin.Context) {
	year := c.Param("year")
	formatDate := time.Now().Format("2006-01-02")
	transactions, err := h.reportUC.GetYearTransactionReport(year)
	if err != nil {
		json.NewResponseError(c, err.Error(), "03", "02")
		return
	}

	mostCommonRoomName, err := h.reportUC.GetMostFrequentRoomNameYears(year)
	if err != nil {
		json.NewResponseError(c, err.Error(), "03", "03")
		return
	}

	file := xlsx.NewFile()

	sheetName := "Year Transactions " + formatDate
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		json.NewResponseError(c, err.Error(), "03", "04")
		return
	}

	headers := []string{"Transaction_id", "FullName", "Room Name", "StartDate", "Description", "EndDate", "Approval Status", "Approved BY", "Created_at", "Updated_at"}
	headerRow := sheet.AddRow()
	for _, header := range headers {
		cell := headerRow.AddCell()
		cell.SetString(header)
	}

	totalTransactions := 0

	for _, transaction := range transactions {
		dataRow := sheet.AddRow()
		dataRow.AddCell().SetString(transaction.Transaction_id)
		dataRow.AddCell().SetString(transaction.Employee_id)
		dataRow.AddCell().SetString(transaction.Room_id)
		dataRow.AddCell().SetString(transaction.StartDate.Format("2006-01-02"))
		dataRow.AddCell().SetString(transaction.Description)
		dataRow.AddCell().SetString(transaction.EndDate.Format("2006-01-02"))
		dataRow.AddCell().SetString(transaction.Approval_status)
		dataRow.AddCell().SetString(transaction.Approved_by)
		dataRow.AddCell().SetString(transaction.Created_at.Format("2006-01-02"))
		dataRow.AddCell().SetString(transaction.Updated_at.Format("2006-01-02"))

		totalTransactions++
	}

	totalRow := sheet.AddRow()
	totalRow.AddCell().SetString("TOTAL TRANSACTION")
	totalRow.AddCell().SetString(fmt.Sprintf("%d", totalTransactions))

	mostCommonRoomRow := sheet.AddRow()
	mostCommonRoomRow.AddCell().SetString("COMMON ROOM")
	mostCommonRoomRow.AddCell().SetString(mostCommonRoomName)

	excelDirectory := "./.excel"
	excelFilename := "Year_transactions_" + formatDate + ".xlsx"
	excelFilePath := filepath.Join(excelDirectory, excelFilename)
	if err := file.Save(excelFilePath); err != nil {
		json.NewResponseError(c, err.Error(), "03", "05")
		return
	}

	if err != nil {
		json.NewResponseError(c, err.Error(), "03", "06")
		return
	}

	json.NewResponseSuccess(c, nil, nil, "Data Berhasil Di Export", "03", "02")

}
