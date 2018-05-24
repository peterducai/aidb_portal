package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//GetLocations get locations
func GetLocations(c echo.Context) error {

	locs, err := AllLocations()
	var tabl = "<table>"
	var tabhead = `<tr>
    <th>ID</th>
    <th>Description</th>
	<th>Country</th>
	<th>Town</th>
	<th>Street</th>
  </tr>`

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "issue with DB connection")
	}

	tabl = tabl + tabhead

	for _, lc := range locs {
		fmt.Println(lc.id, lc.description, lc.country, lc.town, lc.street)
		rw := fmt.Sprintln("<tr><td>", lc.id, "</td><td>", lc.description, "</td><td>", lc.country, "</td><td>", lc.town, "</td><td>", lc.street, "</td>")
		tabl = tabl + rw
	}
	tabl = tabl + "</tr></table>"

	return c.HTML(http.StatusOK, HEADER+tabl+FOOTER)
}
