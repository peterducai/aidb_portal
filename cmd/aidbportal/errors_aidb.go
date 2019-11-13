package aidbportal

import "log"

//ErrCheck check err
func ErrCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
