package router

import (
	"BookingRoom/src/auth/authDelivery"
	"BookingRoom/src/auth/authRepository"
	"BookingRoom/src/auth/authUsecase"
	"BookingRoom/src/employees/employeeDelivery"
	"BookingRoom/src/employees/employeeRepository"
	"BookingRoom/src/employees/employeeUsecase"
	"BookingRoom/src/transactions/transactionsDelivery"
	"BookingRoom/src/transactions/transactionsRepository"
	"BookingRoom/src/transactions/transactionsUseCase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	authRepo := authRepository.NewAuthRepository(db)
	employeeRepo := employeeRepository.NewEmployeeRepository(db)
	transactionsRepo := transactionsRepository.NewTransactionsRepository(db)

	// usecase
	authUC := authUsecase.NewAuthUsecase(authRepo)
	employeeUC := employeeUsecase.NewEmployeeUsecase(employeeRepo)
	transactionUC := transactionsUseCase.NewTransactionsUseCase(transactionsRepo, employeeRepo)

	// delivery
	authDelivery.NewAuthDelivery(v1Group, authUC)
	employeeDelivery.NewEmployeeDelivery(v1Group, employeeUC)
	transactionsDelivery.NewTransactionsDelivery(v1Group, transactionUC)
}
