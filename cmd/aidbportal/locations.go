package aidbportal

import (
	"database/sql"
	"fmt"
)

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

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// Make a slice for the values
	values := make([]interface{}, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	} //-------------------

	locs := make([]*Location, 0)
	for rows.Next() {
		lc := new(Location)
		err := rows.Scan(&lc.id, &lc.description, &lc.country, &lc.town, &lc.street, &lc.zipcode)
		if err != nil {
			return nil, err
		}

		//---------------------
		// Print data
		for i, value := range values {
			switch value.(type) {
			case nil:
				fmt.Println(columns[i], ": NULL")

			case []byte:
				fmt.Println(columns[i], ": ", string(value.([]byte)))

			default:
				fmt.Println(columns[i], ": ", value)
			}
		}
		//---------------------

		locs = append(locs, lc)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return locs, nil
}
