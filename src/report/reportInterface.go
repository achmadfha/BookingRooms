package Report

import (
	"BookingRoom/model/dto/reportDto"
)

type ReportRepository interface {
	GetDailyTransactions(year, month, day string) ([]reportDto.Transactions, error)
	GetDailyTransactionsReport(year, month, day string) ([]reportDto.Transactions, error)
	GetMostFrequentRoomNameDay(year, month, day string) (string, error)
	GetMonthlyTransactions(year, month string) ([]reportDto.Transactions, error)
	GetMonthlyTransactionsReport(year, month string) ([]reportDto.Transactions, error)
	GetMostFrequentRoomNameMonth(year, month string) (string, error)
	GetMostFrequentRoomNameYear(year string) (string, error)
	GetYearTransactions(year string) ([]reportDto.Transactions, error)
	GetYearTransactionsReport(year string) ([]reportDto.Transactions, error)
}

type ReportUsecase interface {
	GetDailyTransactionReport(year, month, day string) ([]reportDto.Transactions, error)
	GetDailyTransaction(year, month, day string) ([]reportDto.Transactions, error)
	GetMostFrequentRoomNamesDay(year, month, day string) (string, error)
	GetMonthlyTransaction(year, month string) ([]reportDto.Transactions, error)
	GetMonthlyTransactionReport(year, month string) ([]reportDto.Transactions, error)
	GetMostFrequentRoomNameMonths(year, month string) (string, error)
	GetYearTransaction(year string) ([]reportDto.Transactions, error)
	GetMostFrequentRoomNameYears(year string) (string, error)
	GetYearTransactionReport(year string) ([]reportDto.Transactions, error)
}
