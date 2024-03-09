package roomsDto

import (
	"github.com/google/uuid"
)

type (
	Rooms struct {
		ID           uuid.UUID `json:"room_id"`
		RoomDetailID uuid.UUID `json:"room_details_id"`
		Name         string    `json:"name"`
		Status       string    `json:"status"`
	}

	RoomResponse struct {
		ID           uuid.UUID   `json:"room_id"`
		RoomDetailID RoomDetails `json:"room_details"`
		Name         string      `json:"name"`
		Status       string      `json:"status"`
	}

	RoomDetails struct {
		RoomDetailsID uuid.UUID `json:"room_details_id"`
		RoomType      string    `json:"room_type"`
		Capacity      int       `json:"capacity"`
		Facility      []string  `json:"facility"`
	}

	RoomsRequest struct {
		Name     string   `json:"name"`
		Status   string   `json:"status"`
		RoomType string   `json:"room_type"`
		Capacity int      `json:"capacity"`
		Facility []string `json:"facility"`
	}

	RoomsCreate struct {
		RoomID        uuid.UUID `json:"room_id"`
		RoomDetailsID uuid.UUID `json:"room_details_id"`
		Name          string    `json:"name"`
		Status        string    `json:"status"`
		RoomType      string    `json:"room_type"`
		Capacity      int       `json:"capacity"`
		Facility      []string  `json:"facility"`
	}
)
