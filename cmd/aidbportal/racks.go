package aidbportal

import "database/sql"

//Rack in the data center
type Rack struct {
	id            sql.NullInt64
	rackName      sql.NullString
	rackSize      sql.NullInt64
	purchaseOrder sql.NullInt64
	rack          sql.NullString
}

//AllRacks will return Racks
func AllRacks() ([]*Rack, error) {
	rows, err := db.Query("SELECT * FROM aidb.rack")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	Racks := make([]*Rack, 0)
	for rows.Next() {
		rk := new(Rack)
		err := rows.Scan(&rk.id, &rk.rackName, &rk.rackSize, &rk.purchaseOrder, &rk.rack)
		if err != nil {
			return nil, err
		}
		Racks = append(Racks, rk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return Racks, nil
}
