package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//CreateJump get locations
func CreateJump(c echo.Context) error {

	locs, err := AllJumps()

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "issue with DB connection")
	}

	//create new html table and populate it with data
	var mytab HtmlTable
	mytab.header = "<table><tr>" //here append <th> tags and </tr> which is table header
	mytab.rows = "<tr><td>"      //here append <td> tags and </td></tr> which is table data
	mytab.footer = "</table>"

	for _, lc := range locs {
		//rw := fmt.Sprintln("</td><td>", lc.description, "</td><td>", lc.country, "</td><td>", lc.town, "</td><td>", lc.street, "</td>")
		//tabl = tabl + rw
		fmt.Println(lc.town)
	}

	return c.HTML(http.StatusOK, HEADER+mytab.header+mytab.rows+mytab.footer+FOOTER)
}
