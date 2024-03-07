package authDelivery

import (
	"BookingRoom/model/dto/employeesDto"
	"BookingRoom/model/dto/json"
	"BookingRoom/pkg/middleware"
	"BookingRoom/pkg/utils"
	"BookingRoom/src/auth"

	"github.com/gin-gonic/gin"
)

type authDelivery struct {
	authUC auth.AuthUsecase
}

func NewAuthDelivery(v1Group *gin.RouterGroup, authUC auth.AuthUsecase) {
	handler := &authDelivery{
		authUC: authUC,
	}
	authGroup := v1Group.Group("/auth")
	authGroup.POST("/login", middleware.BasicAuth, handler.getLogin)
	authGroup.POST("/password", middleware.JWTAuth("ADMIN", "GA", "EMPLOYEE"), handler.setPassword)
}

func (e *authDelivery) getLogin(ctx *gin.Context) {
	var req employeesDto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), "02", "02")
		return
	}

	valErr := utils.ValidationLogin(req)
	if len(valErr) > 0 {
		json.NewResponseBadRequest(ctx, valErr, "Bad Request", "01", "01")
		return
	}

	token, err := e.authUC.Login(req)
	if err != nil {
		switch err.Error() {
		case "01":
			json.NewResponseBadRequest(ctx, nil, "employee doesn't exists on our record", "01", "01")
			return
		case "02":
			json.NewResponseUnauthorized(ctx, "Unauthorized: Username and password do not match", "02", "02")
			return
		case "03":
			json.NewResponseUnauthorized(ctx, "Invalid Token Access", "03", "03")
			return
		default:
			json.NewResponseError(ctx, err.Error(), "04", "04")
			return
		}
	}

	data := interface{}(map[string]interface{}{"access_token": token})

	json.NewResponseSuccess(ctx, data, nil, "success", "01", "03")
}

func (e *authDelivery) setPassword(ctx *gin.Context) {
	var req employeesDto.PasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), "02", "02")
		return
	}

	if err := e.authUC.UpdatePassword(req); err != nil {
		switch err.Error() {
		case "01":
			json.NewResponseBadRequest(ctx, nil, "Employee doesn't exist in our records", "01", "01")
		case "02":
			json.NewResponseBadRequest(ctx, nil, "Old password is incorrect", "02", "02")
		case "03":
			json.NewResponseBadRequest(ctx, nil, "New password and confirm password do not match", "03", "03")
		case "04":
			json.NewResponseBadRequest(ctx, nil, "Internal Server Error", "04", "04")
		case "05":
			json.NewResponseError(ctx, "Failed to update password", "05", "05")
		default:
			json.NewResponseError(ctx, "Unknown error occurred", "06", "06")
		}
		return
	}

	json.NewResponseSuccess(ctx, nil, nil, "Password updated successfully.", "01", "01")
}
