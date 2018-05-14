package main

import (
	"database/sql"
	"fmt"
	"log"
)

type server struct {
	db *sql.DB
}

//DoQuery qe
func DoQuery(db *sql.DB) ([]string, error) {

	var data []string
	rows, err := db.Query("SELECT * FROM aidb.country")
	ErrCheck(err)
	defer rows.Close()

	for {
		// cols, err := rows.Columns()
		// if err != nil {
		// 	return err
		// }
		for rows.Next() {
			var (
				id        sql.NullInt64
				iso       sql.NullString
				wholename sql.NullString
				nicename  sql.NullString
				iso3      sql.NullString
				numcode   sql.NullInt64
				phonecode sql.NullInt64
			)
			err := rows.Scan(&id, &iso, &wholename, &nicename, &iso3, &numcode, &phonecode)
			ErrCheck(err)
			fmt.Printf("id %d name is %s and zip is %d\n", id, nicename, numcode)
		}
		if err := rows.Err(); err != nil {
			return data, err
		}
		if !rows.NextResultSet() {
			return data, nil
		}

		return data, nil
	}
	log.Fatal("unreachable")
	// var roleMap = map[int64]string{
	// 	1: "user",
	// 	2: "admin",
	// 	3: "gopher",
	// }
	// for rows.Next() {
	// 	var (
	// 		id   int64
	// 		role int64
	// 	)
	// 	if err := rows.Scan(&id, &role); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(id)
	// 	//fmt.Printf("id %d has role %s\n", id, roleMap[role])
	// }
	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	return data, nil
}
