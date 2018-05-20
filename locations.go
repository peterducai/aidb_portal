package main

import "database/sql"

//Location of data center
type Location struct {
	id          sql.NullInt64
	description sql.NullString
	country     sql.NullString
	town        sql.NullString
	street      sql.NullString
	zipcode     sql.NullString
}

//AllLocations will return locations
func AllLocations() ([]*Location, error) {
	rows, err := db.Query("SELECT * FROM aidb.location")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	locs := make([]*Location, 0)
	for rows.Next() {
		lc := new(Location)
		err := rows.Scan(&lc.id, &lc.description, &lc.country, &lc.town, &lc.street, &lc.zipcode)
		if err != nil {
			return nil, err
		}
		locs = append(locs, lc)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return locs, nil
}
