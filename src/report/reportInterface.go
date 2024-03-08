package Report

import (
	"BookingRoom/model/dto/reportDto"
)

type ReportRepository interface {
	GetDailyTransactions(created_at string) ([]reportDto.Transactions, error)
	GetMonthlyTransactions(year, month string) ([]reportDto.Transactions, error)
	GetYearTransactions(year string) ([]reportDto.Transactions, error)
}

type ReportUsecase interface {
	GetDailyTransaction(created_at string) ([]reportDto.Transactions, error)
	GetMonthlyTransaction(year, month string) ([]reportDto.Transactions, error)
	GetYearTransaction(year string) ([]reportDto.Transactions, error)
}
