package aidbportal

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

//InitDB initialize db
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

// //FormatResults will change data to string (for html table)
// func FormatResults(rows QueryRow) {

// }
