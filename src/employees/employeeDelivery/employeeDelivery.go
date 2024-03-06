package employeeDelivery

import (
	"BookingRoom/model/dto"
	"BookingRoom/model/dto/json"
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
		employeeGroup.GET("/", handler.getEmployee)
		employeeGroup.GET("/:id", handler.getEmployeeById)
		employeeGroup.POST("/", handler.createEmployee)
		employeeGroup.PUT("/:id", handler.updateEmployee)
		employeeGroup.DELETE("/:id", handler.deleteEmployee)
	}
}

func (e *employeeDelivery) getEmployee(ctx *gin.Context) {
	employee, err := e.employeeUC.GetEmployee()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "02", "02")
		return
	}

	json.NewResponseSuccess(ctx, employee, nil, "success", "01", "01")
}

func (e *employeeDelivery) getEmployeeById(ctx *gin.Context) {
	employeeId := ctx.Param("id")

	employee, err := e.employeeUC.GetEmployeeById(employeeId)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, employee, nil, "success", "01", "01")
}

func (e *employeeDelivery) createEmployee(ctx *gin.Context) {
	var employee dto.Employees

	if err := ctx.ShouldBindJSON(&employee); err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, err.Error(), "01", "01")
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

	var employee dto.Employees
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, err.Error(), "01", "01")
		return
	}

	employee.EmployeeId = employeeId

	if err := e.employeeUC.UpdateEmployee(employee); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, employee, nil, "success", "01", "01")
}

func (e *employeeDelivery) deleteEmployee(ctx *gin.Context) {
	employeeId := ctx.Param("id")

	if err := e.employeeUC.DeleteEmployeeById(employeeId); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, "OK", nil, "success", "01", "01")
}
