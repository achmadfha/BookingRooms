package authDelivery

import (
	"BookingRoom/model/dto"
	"BookingRoom/model/dto/json"
	"BookingRoom/pkg/middleware"
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
}

func (e *authDelivery) getLogin(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		detail := json.ValidationField{FieldName: "Login", Message: err.Error()}
		listError := []json.ValidationField{detail}
		json.NewResponseBadRequest(ctx, listError, "Bad Request", "01", "01")
		return
	}

	token, err := e.authUC.Login(req)
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
