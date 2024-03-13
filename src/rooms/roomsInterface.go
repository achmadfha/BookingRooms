package rooms

import (
	"BookingRoom/model/dto/json"
	"BookingRoom/model/dto/roomsDto"
)

type RoomsRepository interface {
	CreateRooms(room roomsDto.RoomsCreate) error
	RetrieveAllRooms(page, pageSize int) ([]roomsDto.Rooms, error)
	CountAllRooms() (int, error)
	RetrieveRoomsByID(roomID string) (roomsDto.RoomResponse, error)
	UpdateRooms(room roomsDto.RoomsCreate) error
}

type RoomUseCase interface {
	CreateRooms(req roomsDto.RoomsRequest) (roomsDto.RoomResponse, error)
	RetrieveAllRooms(page, pageSize int) ([]roomsDto.Rooms, json.Pagination, error)
	RetrieveRoomByID(roomID string) (roomsDto.RoomResponse, error)
	UpdateRoomsByID(req roomsDto.RoomsCreate) error
}
