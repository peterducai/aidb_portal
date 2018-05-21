package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//GetRooms get locations
func GetRooms(c echo.Context) error {

	rooms, err := AllLocations()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "333")
	}
	for _, rm := range rooms {
		fmt.Println(rm.id, rm.description, rm.country, rm.town, rm.street)
	}

	return c.String(http.StatusOK, "Rooms!")
}
