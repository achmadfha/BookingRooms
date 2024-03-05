package router

import (
	"BookingRoom/src/employees/employeeDelivery"
	"BookingRoom/src/employees/employeeRepository"
	"BookingRoom/src/employees/employeeUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRouter(v1Group *gin.RouterGroup, db *sql.DB) {
	// repository
	employeeRepo := employeeRepository.NewEmployeeRepository(db)

	// usecase
	employeeUC := employeeUsecase.NewEmployeeUsecase(employeeRepo)

	// delivery
	employeeDelivery.NewEmployeeDelivery(v1Group, employeeUC)
}
