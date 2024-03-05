package employeeDelivery

import (
	"BookingRoom/model/dto"
	"BookingRoom/model/dto/json"
	"BookingRoom/pkg/middleware"
	"BookingRoom/pkg/utils"
	"BookingRoom/src/employees"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	var req dto.Employees
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, err.Error(), "01", "01")
		return
	}

	employee, err := e.employeeUC.GetLogin(req.Username)
	if err != nil {
		fmt.Println("err Delivery :", err.Error())
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, "Invalid username or password", "01", "01")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(req.Password))
	if err != nil {
		json.NewResponseBadRequest(ctx, []json.ValidationField{}, "Invalid hash password", "01", "01")
		return
	}

	token, err := utils.GenerateToken(employee.EmployeeId, employee.Position)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
	}

	json.NewResponseSuccess(ctx, token, nil, "success", "01", "01")
}
