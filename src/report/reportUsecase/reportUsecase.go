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

func (t *reportUC) GetDailyTransaction(created_at string) ([]reportDto.Transactions, error) {
	return t.reportRepo.GetDailyTransactions(created_at)
}

func (t *reportUC) GetMonthlyTransaction(year, month string) ([]reportDto.Transactions, error) {
	return t.reportRepo.GetMonthlyTransactions(year, month)
}

func (t *reportUC) GetYearTransaction(year string) ([]reportDto.Transactions, error) {
	return t.reportRepo.GetYearTransactions(year)
}
