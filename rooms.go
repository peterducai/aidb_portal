package main

import "database/sql"

//Room of data center
type Room struct {
	id         sql.NullInt64
	loc        sql.NullInt64
	floor      sql.NullString
	roomNumber sql.NullInt64
	roomName   sql.NullString
}

//AllRooms will return Rooms
func AllRooms() ([]*Room, error) {
	rows, err := db.Query("SELECT * FROM aidb.room")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := make([]*Room, 0)
	for rows.Next() {
		rm := new(Room)
		err := rows.Scan(&rm.id, &rm.loc, &rm.floor, &rm.roomNumber, &rm.roomName)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, rm)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}
