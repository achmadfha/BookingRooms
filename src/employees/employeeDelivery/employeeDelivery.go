package employeeDelivery

import (
	"BookingRoom/model/dto/employeesDto"
	"BookingRoom/model/dto/json"
	"BookingRoom/pkg/middleware"
	"BookingRoom/src/employees"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type employeeDelivery struct {
	employeeUC employees.EmployeeUsecase
}

func NewEmployeeDelivery(v1Group *gin.RouterGroup, employeeUC employees.EmployeeUsecase) {
	handler := &employeeDelivery{
		employeeUC: employeeUC,
	}

	employeeGroup := v1Group.Group("/employees")
	{
		employeeGroup.GET("/", middleware.JWTAuth("ADMIN", "GA"), handler.getEmployee)
		employeeGroup.GET("/:id", middleware.JWTAuth("ADMIN", "GA", "EMPLOYEE"), handler.getEmployeeById)
		employeeGroup.POST("/", middleware.JWTAuth("ADMIN", "GA"), handler.createEmployee)
		employeeGroup.PUT("/:id", middleware.JWTAuth("ADMIN", "GA", "EMPLOYEE"), handler.updateEmployee)
		employeeGroup.DELETE("/:id", middleware.JWTAuth("ADMIN", "GA"), handler.deleteEmployee)
	}
}

func (e *employeeDelivery) getEmployee(ctx *gin.Context) {
	page := ctx.Query("page")
	size := ctx.Query("size")
	employee, pagination, err := e.employeeUC.GetEmployee(page, size)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseSuccess(ctx, "Data Not Found", nil, "success", "01", "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), "02", "02")
		return
	}

	json.NewResponseSuccess(ctx, employee, pagination, "success", "01", "01")
}

func (e *employeeDelivery) getEmployeeById(ctx *gin.Context) {
	employeeId := ctx.Param("id")

	employee, err := e.employeeUC.GetEmployeeById(employeeId)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseSuccess(ctx, "Data Not Found", nil, "success", "01", "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), "02", "02")
		return
	}

	json.NewResponseSuccess(ctx, employee, nil, "success", "01", "01")
}

func (e *employeeDelivery) createEmployee(ctx *gin.Context) {
	var employee employeesDto.Employees

	if err := ctx.ShouldBindJSON(&employee); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	valError := employeesDto.ValidationEmployee(employee)
	if len(valError) > 0 {
		json.NewResponseBadRequest(ctx, valError, "failed", "01", "01")
		return
	}

	if err := e.employeeUC.StoreEmployee(&employee); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, employee, nil, "success", "01", "01")
}

func (e *employeeDelivery) updateEmployee(ctx *gin.Context) {
	employeeIdString := ctx.Param("id")
	employeeId, err := uuid.Parse(employeeIdString)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	var employee employeesDto.Employees
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	valError := employeesDto.ValidationEmployee(employee)
	if len(valError) > 0 {
		json.NewResponseBadRequest(ctx, valError, "failed", "01", "01")
		return
	}

	employee.EmployeeId = employeeId

	if err := e.employeeUC.UpdateEmployee(employee); err != nil {
		if err.Error() == "1" {
			json.NewResponseError(ctx, "Id not found", "02", "02")
			return
		}
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, employee, nil, "success", "01", "01")
}

func (e *employeeDelivery) deleteEmployee(ctx *gin.Context) {
	employeeId := ctx.Param("id")

	if err := e.employeeUC.DeleteEmployeeById(employeeId); err != nil {
		if err.Error() == "1" {
			json.NewResponseError(ctx, "Id not found", "02", "02")
			return
		}
		json.NewResponseError(ctx, err.Error(), "02", "02")
		return
	}

	json.NewResponseSuccess(ctx, "OK", nil, "success", "01", "01")
}
