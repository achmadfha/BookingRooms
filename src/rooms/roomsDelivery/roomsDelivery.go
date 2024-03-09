package roomsDelivery

import (
	"BookingRoom/model/dto/json"
	"BookingRoom/model/dto/roomsDto"
	"BookingRoom/pkg/middleware"
	"BookingRoom/pkg/utils"
	"BookingRoom/src/rooms"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
)

type roomsDelivery struct {
	roomsUC rooms.RoomUseCase
}

func NewRoomDelivery(v1Group *gin.RouterGroup, roomsUC rooms.RoomUseCase) {
	handler := roomsDelivery{roomsUC: roomsUC}

	roomsGroup := v1Group.Group("/room")
	{
		roomsGroup.POST("", middleware.JWTAuth("ADMIN", "GA"), handler.CreateRooms)
		roomsGroup.GET("", middleware.JWTAuth("ADMIN", "EMPLOYEE", "GA"), handler.RetrieveAllRooms)
		roomsGroup.GET("/:id", middleware.JWTAuth("ADMIN", "EMPLOYEE", "GA"), handler.RetrieveRoomsByID)
		roomsGroup.PUT("/:id", middleware.JWTAuth("ADMIN", "GA"), handler.UpdateRooms)
	}
}

func (r roomsDelivery) CreateRooms(ctx *gin.Context) {
	var req roomsDto.RoomsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), "02", "01")
		return
	}

	validationError := utils.RoomsValidation(req)
	if len(validationError) > 0 {
		json.NewResponseBadRequest(ctx, validationError, "failed", "02", "01")
		return
	}

	err := r.roomsUC.CreateRooms(req)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseError(ctx, "err", "02", "01")
			return
		}

		if err.Error() == "02" {
			json.NewResponseError(ctx, "err", "02", "01")
			return
		}

		json.NewResponseError(ctx, err.Error(), "02", "01")
		return
	}

	json.NewResponseCreatedSuccess(ctx, "Rooms created successfully", "02", "01")
}

func (r roomsDelivery) RetrieveAllRooms(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("size"))

	data, pagination, err := r.roomsUC.RetrieveAllRooms(page, pageSize)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseError(ctx, "page not found", "02", "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), "02", "02")
		return
	}

	json.NewResponseSuccess(ctx, data, pagination, "success", "02", "02")
}

func (r roomsDelivery) RetrieveRoomsByID(ctx *gin.Context) {
	roomID := ctx.Param("id")

	data, err := r.roomsUC.RetrieveRoomByID(roomID)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseBadRequest(ctx, nil, "Transaction ID not found on our records", "02", "03")
			return
		}
		json.NewResponseError(ctx, err.Error(), "02", "03")
		return
	}

	json.NewResponseSuccess(ctx, data, nil, "success", "02", "03")
}

func (r roomsDelivery) UpdateRooms(ctx *gin.Context) {
	var req roomsDto.RoomsCreate
	roomIDStr := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseError(ctx, err.Error(), "03", "04")
		return
	}

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "02", "04")
		return
	}

	roomsData := roomsDto.RoomsCreate{
		RoomID:   roomID,
		Name:     req.Name,
		Status:   req.Status,
		RoomType: req.RoomType,
		Capacity: req.Capacity,
		Facility: req.Facility,
	}

	err = r.roomsUC.UpdateRoomsByID(roomsData)
	if err != nil {
		if err.Error() == "01" {
			json.NewResponseBadRequest(ctx, nil, "Rooms ID not found on our records", "02", "04")
			return
		}

		if err.Error() == "02" {
			json.NewResponseBadRequest(ctx, nil, "Status status must be either AVAILABLE or BOOKED", "02", "04")
			return
		}

		json.NewResponseError(ctx, err.Error(), "02", "04")
		return
	}

	json.NewResponseSuccess(ctx, "Transaction log updated successfully", nil, "success", "02", "04")
}
