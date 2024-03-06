package router

import (
	"BookingRoom/src/employees/employeeDelivery"
	"BookingRoom/src/employees/employeeRepository"
	"BookingRoom/src/employees/employeeUsecase"
	"BookingRoom/src/transactions/transactionDelivery"
	"BookingRoom/src/transactions/transactionRepository"
	"BookingRoom/src/transactions/transactionUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	employeeRepo := employeeRepository.NewEmployeeRepository(db)
	transactionRepo := transactionRepository.NewTransactionRepo(db)

	// usecase
	employeeUC := employeeUsecase.NewEmployeeUsecase(employeeRepo)
	transactionUC := transactionUsecase.NewTransactionUsecase(transactionRepo)

	// delivery
	employeeDelivery.NewEmployeeDelivery(v1Group, employeeUC)
	transactionDelivery.NewTransactionDelivery(v1Group, transactionUC)
}
