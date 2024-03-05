package employeeDelivery

import (
	"BookingRoom/model/dto"
	"BookingRoom/model/dto/json"
	"BookingRoom/pkg/middleware"
	"BookingRoom/src/employees"

	"github.com/gin-gonic/gin"
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
		employeeGroup.POST("/login", middleware.BasicAuth, handler.getLogin)
		employeeGroup.GET("/", middleware.JWTAuth("ADMIN", "GA"))
		employeeGroup.POST("/", middleware.JWTAuth("ADMIN", "GA"))
		employeeGroup.GET("/:id", middleware.JWTAuth("ADMIN", "GA"))
		employeeGroup.PUT("/:id", middleware.JWTAuth("ADMIN", "GA"))
		employeeGroup.DELETE("/:id", middleware.JWTAuth("GA"))
	}
}

func (e *employeeDelivery) getLogin(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		detail := json.ValidationField{FieldName: "Login", Message: err.Error()}
		listError := []json.ValidationField{detail}
		json.NewResponseBadRequest(ctx, listError, "Bad Request", "01", "01")
		return
	}

	token, err := e.employeeUC.Login(req)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseBadRequest(ctx, nil, "employee doesn't exists on our record", "01", "02")
			return
		}
		if err.Error() == "02" {
			json.NewResponseBadRequest(ctx, nil, "Unauthorized username and password didn't match", "01", "02")
			return
		}
		json.NewResponseError(ctx, err.Error(), "01", "02")
		return
	}

	data := interface{}(map[string]interface{}{"access_token": token})

	json.NewResponseSuccess(ctx, data, nil, "success", "01", "01")
}
