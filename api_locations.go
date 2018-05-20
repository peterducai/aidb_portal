package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//GetLocations get locations
func GetLocations(c echo.Context) error {

	locs, err := AllLocations()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "333")
	}
	for _, lc := range locs {
		fmt.Println(lc.id, lc.description, lc.country, lc.town, lc.street)
	}

	return c.String(http.StatusOK, "Locations!")
}
