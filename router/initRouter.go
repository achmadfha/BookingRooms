package router

import (
	"BookingRoom/src/employees/employeeDelivery"
	"BookingRoom/src/employees/employeeRepository"
	"BookingRoom/src/employees/employeeUsecase"
	"BookingRoom/src/report/reportDelivery"
	"BookingRoom/src/report/reportRepository"
	"BookingRoom/src/report/reportUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	employeeRepo := employeeRepository.NewEmployeeRepository(db)
	//transactionRepo := transactionRepository.NewTransactionRepo(db)
	reportRepo := reportRepository.NewReportRepo(db)

	// usecase
	employeeUC := employeeUsecase.NewEmployeeUsecase(employeeRepo)
	//transactionUC := transactionUsecase.NewTransactionUsecase(transactionRepo)
	reportUC := reportUsecase.NewReportUsecase(reportRepo)

	// delivery
	employeeDelivery.NewEmployeeDelivery(v1Group, employeeUC)
	//transactionDelivery.NewTransactionDelivery(v1Group, transactionUC)
	reportDelivery.NewReportDelivery(v1Group, reportUC)
}
