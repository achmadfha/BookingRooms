package router

import (
	"BookingRoom/src/auth/authDelivery"
	"BookingRoom/src/auth/authRepository"
	"BookingRoom/src/auth/authUsecase"
	"BookingRoom/src/employees/employeeDelivery"
	"BookingRoom/src/employees/employeeRepository"
	"BookingRoom/src/employees/employeeUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	authRepo := authRepository.NewAuthRepository(db)
	employeeRepo := employeeRepository.NewEmployeeRepository(db)

	// usecase
	authUC := authUsecase.NewAuthUsecase(authRepo)
	employeeUC := employeeUsecase.NewEmployeeUsecase(employeeRepo)

	// delivery
	authDelivery.NewAuthDelivery(v1Group, authUC)
	employeeDelivery.NewEmployeeDelivery(v1Group, employeeUC)
}
