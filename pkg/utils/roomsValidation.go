package utils

import (
	"BookingRoom/model/dto/json"
	"BookingRoom/model/dto/roomsDto"
)

func RoomsValidation(req roomsDto.RoomsRequest) []json.ValidationField {
	var validationErrors []json.ValidationField

	// Name validation
	if req.Name == "" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "name",
			Message:   "Name is required",
		})
	} else if len(req.Name) < 5 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "name",
			Message:   "Name should be at least 5 characters",
		})
	}

	// Status validation
	if req.Status != "AVAILABLE" && req.Status != "BOOKED" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "status",
			Message:   "Status status must be either AVAILABLE or BOOKED",
		})
	}

	// RoomType validation
	if req.RoomType == "" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "room_type",
			Message:   "Room Type is required",
		})
	} else if len(req.RoomType) < 3 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "room_type",
			Message:   "Room Type be at least 3 characters",
		})
	}

	// Capacity validation
	if req.Capacity <= 0 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "capacity",
			Message:   "Capacity should be greater than 0",
		})
	}

	// Facility validation
	if len(req.Facility) == 0 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "facility",
			Message:   "At least one facility is required",
		})
	}

	return validationErrors
}
