package Report

import (
	"BookingRoom/model/dto/reportDto"
)

type ReportRepository interface {
	GetDailyTransactions(year, month, day string) ([]reportDto.Transactions, error)
	GetDailyTransactionsReport(year, month, day string) ([]reportDto.Transactions, error)
	GetMonthlyTransactions(year, month string) ([]reportDto.Transactions, error)
	GetMonthlyTransactionsReport(year, month string) ([]reportDto.Transactions, error)
	GetYearTransactions(year string) ([]reportDto.Transactions, error)
	GetYearTransactionsReport(year string) ([]reportDto.Transactions, error)
}

type ReportUsecase interface {
	GetDailyTransactionReport(year, month, day string) ([]reportDto.Transactions, error)
	GetDailyTransaction(year, month, day string) ([]reportDto.Transactions, error)
	GetMonthlyTransaction(year, month string) ([]reportDto.Transactions, error)
	GetMonthlyTransactionReport(year, month string) ([]reportDto.Transactions, error)
	GetYearTransaction(year string) ([]reportDto.Transactions, error)
	GetYearTransactionReport(year string) ([]reportDto.Transactions, error)
}
