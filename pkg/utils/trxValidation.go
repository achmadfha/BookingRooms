package utils

import (
	"BookingRoom/model/dto/json"
	"BookingRoom/model/dto/transactionsDto"
	"github.com/google/uuid"
	"time"
)

func ValidationTrxReq(trxReq transactionsDto.TransactionsRequest) []json.ValidationField {
	var validationErrors []json.ValidationField

	// Validate EmployeeId
	if trxReq.EmployeeId == uuid.Nil {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "employee_id",
			Message:   "Employee ID is required",
		})
	}

	// Validate RoomId
	if trxReq.RoomId == uuid.Nil {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "room_id",
			Message:   "Room ID is required",
		})
	}

	if trxReq.StartDate == "" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "start_date",
			Message:   "Start date is required",
		})
	} else {
		_, err := time.Parse("2006-01-02", trxReq.StartDate)
		if err != nil {
			validationErrors = append(validationErrors, json.ValidationField{
				FieldName: "start_date",
				Message:   "Start date should be in the format YYYY-MM-DD",
			})
		}
	}

	if trxReq.EndDate == "" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "end_date",
			Message:   "End date is required",
		})
	} else {
		_, err := time.Parse("2006-01-02", trxReq.EndDate)
		if err != nil {
			validationErrors = append(validationErrors, json.ValidationField{
				FieldName: "end_date",
				Message:   "End date should be in the format YYYY-MM-DD",
			})
		}
	}

	if trxReq.Description == "" {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "description",
			Message:   "Description is required",
		})
	} else if len(trxReq.Description) < 10 {
		validationErrors = append(validationErrors, json.ValidationField{
			FieldName: "description",
			Message:   "Description length should be 10 characters minimum",
		})
	}

	return validationErrors
}
