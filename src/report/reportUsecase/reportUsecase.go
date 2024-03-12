package reportUsecase

import (
	"BookingRoom/model/dto/reportDto"
	Report "BookingRoom/src/report"
)

type reportUC struct {
	reportRepo Report.ReportRepository
}

func NewReportUsecase(reportRepo Report.ReportRepository) Report.ReportUsecase {
	return &reportUC{reportRepo}
}

func (t *reportUC) GetDailyTransactionReport(year, month, day string) ([]reportDto.Transactions, error) {
	return t.reportRepo.GetDailyTransactionsReport(year, month, day)
}

func (t *reportUC) GetMostFrequentRoomNamesDay(year, month, day string) (string, error) {
	mostFrequentRoomName, err := t.reportRepo.GetMostFrequentRoomNameDay(year, month, day)
	if err != nil {
		return "", err
	}

	return mostFrequentRoomName, nil
}

func (t *reportUC) GetDailyTransaction(year, month, day string) ([]reportDto.Transactions, error) {
	return t.reportRepo.GetDailyTransactions(year, month, day)
}

func (t *reportUC) GetMonthlyTransactionReport(year, month string) ([]reportDto.Transactions, error) {
	return t.reportRepo.GetMonthlyTransactionsReport(year, month)
}

func (t *reportUC) GetMostFrequentRoomNameMonths(year, month string) (string, error) {
	mostFrequentRoomName, err := t.reportRepo.GetMostFrequentRoomNameMonth(year, month)
	if err != nil {
		return "", err
	}

	return mostFrequentRoomName, nil
}

func (t *reportUC) GetMonthlyTransaction(year, month string) ([]reportDto.Transactions, error) {
	return t.reportRepo.GetMonthlyTransactions(year, month)
}

func (t *reportUC) GetMostFrequentRoomNameYears(year string) (string, error) {
	mostFrequentRoomName, err := t.reportRepo.GetMostFrequentRoomNameYear(year)
	if err != nil {
		return "", err
	}

	return mostFrequentRoomName, nil
}

func (t *reportUC) GetYearTransaction(year string) ([]reportDto.Transactions, error) {
	return t.reportRepo.GetYearTransactions(year)
}

func (t *reportUC) GetYearTransactionReport(year string) ([]reportDto.Transactions, error) {
	return t.reportRepo.GetYearTransactionsReport(year)
}
