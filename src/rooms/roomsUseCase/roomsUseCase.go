package roomsUseCase

import (
	"BookingRoom/model/dto/json"
	"BookingRoom/model/dto/roomsDto"
	"BookingRoom/src/rooms"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"math"
)

type roomsUC struct {
	roomsRepository rooms.RoomsRepository
}

func NewRoomsUseCase(roomsRepo rooms.RoomsRepository) rooms.RoomUseCase {
	return &roomsUC{roomsRepo}
}

func (r roomsUC) CreateRooms(req roomsDto.RoomsRequest) (roomsDto.RoomResponse, error) {
	roomId, err := uuid.NewRandom()
	if err != nil {
		// 01 error while generate uuid room id
		return roomsDto.RoomResponse{}, errors.New("01")
	}

	roomDetailId, err := uuid.NewRandom()
	if err != nil {
		// 02 error while generate uuid room details id
		return roomsDto.RoomResponse{}, errors.New("02")
	}

	newRooms := roomsDto.RoomsCreate{
		RoomID:        roomId,
		RoomDetailsID: roomDetailId,
		Name:          req.Name,
		Status:        req.Status,
		RoomType:      req.RoomType,
		Capacity:      req.Capacity,
		Facility:      req.Facility,
	}

	err = r.roomsRepository.CreateRooms(newRooms)
	if err != nil {
		return roomsDto.RoomResponse{}, err
	}

	// Construct roomsDetail with desired values
	roomsDetail := roomsDto.RoomDetails{
		RoomDetailsID: roomDetailId,
		RoomType:      req.RoomType,
		Capacity:      req.Capacity,
		Facility:      req.Facility,
	}

	data := roomsDto.RoomResponse{
		ID:           roomId,
		RoomDetailID: roomsDetail,
		Name:         req.Name,
		Status:       req.Status,
	}

	return data, nil
}

func (r roomsUC) RetrieveAllRooms(page, pageSize int) ([]roomsDto.Rooms, json.Pagination, error) {
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 5
	}

	roomsData, err := r.roomsRepository.RetrieveAllRooms(page, pageSize)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, json.Pagination{}, errors.New("no rows found")
		}
		return nil, json.Pagination{}, err
	}

	totalRooms, err := r.roomsRepository.CountAllRooms()
	if err != nil {
		return nil, json.Pagination{}, err
	}

	totalPages := int(math.Ceil(float64(totalRooms) / float64(pageSize)))
	if page > totalPages {
		return nil, json.Pagination{}, errors.New("01")
	}

	if totalPages == 0 && totalRooms > 0 {
		totalPages = 1
	}

	pagination := json.Pagination{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalRecords: totalRooms,
	}

	return roomsData, pagination, nil
}

func (r roomsUC) RetrieveRoomByID(roomID string) (roomsDto.RoomResponse, error) {
	roomData, err := r.roomsRepository.RetrieveRoomsByID(roomID)
	if err != nil {
		if err.Error() == "01" {
			// 01 no rows
			return roomsDto.RoomResponse{}, errors.New("01")
		}
		return roomsDto.RoomResponse{}, err
	}

	return roomData, nil
}

func (r roomsUC) UpdateRoomsByID(req roomsDto.RoomsCreate) error {
	roomId := req.RoomID.String()
	room, err := r.roomsRepository.RetrieveRoomsByID(roomId)
	if err != nil {
		if err.Error() == "01" {
			// no rows
			return errors.New("01")
		}
		return err
	}

	// Validate RoomType
	if req.Status != "" && req.Status != "AVAILABLE" && req.Status != "BOOKED" {
		// invalid RoomType
		return errors.New("02")
	}

	if req.Name != "" {
		room.Name = req.Name
	}
	if req.Status != "" {
		room.Status = req.Status
	}
	if req.RoomType != "" {
		room.RoomDetailID.RoomType = req.RoomType
	}
	if req.Capacity != 0 {
		room.RoomDetailID.Capacity = req.Capacity
	}
	if len(req.Facility) > 0 {
		room.RoomDetailID.Facility = req.Facility
	}

	// how i set if the request contains new values
	roomData := roomsDto.RoomsCreate{
		RoomID:        room.ID,
		RoomDetailsID: room.RoomDetailID.RoomDetailsID,
		Name:          room.Name,
		Status:        room.Status,
		RoomType:      room.RoomDetailID.RoomType,
		Capacity:      room.RoomDetailID.Capacity,
		Facility:      room.RoomDetailID.Facility,
	}

	err = r.roomsRepository.UpdateRooms(roomData)
	if err != nil {
		return err
	}

	return nil
}
