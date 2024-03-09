package roomsRepository

import (
	"BookingRoom/model/dto/roomsDto"
	"BookingRoom/src/rooms"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
)

type roomsRepository struct {
	db *sql.DB
}

func NewRoomsRepository(db *sql.DB) rooms.RoomsRepository {
	return &roomsRepository{db}
}

func (r roomsRepository) CreateRooms(room roomsDto.RoomsCreate) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// insert rooms Details
	facilityArray := pq.Array(room.Facility)
	fmt.Println(facilityArray)
	trxRoomsDetailsQuery := `
		INSERT INTO room_details
			(room_details_id, room_type, capacity, facility)
		VALUES 
		    ($1, $2, $3, $4)`

	_, err = tx.Exec(trxRoomsDetailsQuery, room.RoomDetailsID, room.RoomType, room.Capacity, facilityArray)
	if err != nil {
		return err
	}

	// insert rooms
	trxRoomQuery := `
		INSERT INTO room
		    (room_id, room_details_id, name, status)
		VALUES
		    ($1, $2, $3, $4)`

	_, err = tx.Exec(trxRoomQuery, room.RoomID, room.RoomDetailsID, room.Name, room.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r roomsRepository) RetrieveAllRooms(page, pageSize int) ([]roomsDto.Rooms, error) {
	offset := (page - 1) * pageSize
	limit := pageSize

	query := `SELECT room_id, room_details_id, name, status FROM room LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rooms []roomsDto.Rooms
	for rows.Next() {
		var room roomsDto.Rooms
		err := rows.Scan(&room.ID, &room.RoomDetailID, &room.Name, &room.Status)
		if err != nil {
			errors.New(fmt.Sprintf("Error scanning rooms row: %s", err))
			continue
		}

		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("Error iterating through rooms rows: %s", err))
	}

	return rooms, nil
}

func (r roomsRepository) CountAllRooms() (int, error) {
	var count int

	query := `SELECT COUNT(*) FROM room`

	row := r.db.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r roomsRepository) RetrieveRoomsByID(roomID string) (roomsDto.RoomResponse, error) {
	var roomDetails roomsDto.RoomResponse

	query := `
        SELECT 
            r.room_id,
            r.name,
            r.status,
            rd.room_details_id,
            rd.room_type,
            rd.capacity,
            rd.facility
        FROM 
            room r
        JOIN 
            room_details rd ON r.room_details_id = rd.room_details_id
        WHERE 
            r.room_id = $1
    `

	row := r.db.QueryRow(query, roomID)
	err := row.Scan(
		&roomDetails.ID,
		&roomDetails.Name,
		&roomDetails.Status,
		&roomDetails.RoomDetailID.RoomDetailsID,
		&roomDetails.RoomDetailID.RoomType,
		&roomDetails.RoomDetailID.Capacity,
		pq.Array(&roomDetails.RoomDetailID.Facility),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return roomsDto.RoomResponse{}, errors.New("01")
		}
		return roomsDto.RoomResponse{}, err
	}

	return roomDetails, nil
}

func (r roomsRepository) UpdateRooms(room roomsDto.RoomsCreate) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	// Update room details
	facilityArray := pq.Array(room.Facility)
	trxRoomDetailsQuery := `
        UPDATE room_details
        SET 
            room_type = $2,
            capacity = $3,
            facility = $4
        WHERE
            room_details_id = $1
    `

	_, err = tx.Exec(trxRoomDetailsQuery, room.RoomDetailsID, room.RoomType, room.Capacity, facilityArray)
	if err != nil {
		return err
	}

	// Update room
	trxRoomQuery := `
        UPDATE room
        SET 
            name = $2,
            status = $3
        WHERE
            room_id = $1
    `

	_, err = tx.Exec(trxRoomQuery, room.RoomID, room.Name, room.Status)
	if err != nil {
		return err
	}

	return nil
}
