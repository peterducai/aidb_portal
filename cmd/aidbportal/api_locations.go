package aidbportal

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
		return c.String(http.StatusBadRequest, "issue with DB connection")
	}

	//create new html table and populate it with data

	Page.header = "<table><tr>" //here append <th> tags and </tr> which is table header
	Page.rows = "<tr><td>"      //here append <td> tags and </td></tr> which is table data
	Page.footer = "</table>"

	//for _, lc := range locs {
	//	rw := fmt.Sprintln("</td><td>", lc.description, "</td><td>", lc.country, "</td><td>", lc.town, "</td><td>", lc.street, "</td>")
	//	//tabl = tabl + rw
	//}
	fmt.Println()
	fmt.Println(locs)

	return c.HTML(http.StatusOK, HEADER+Page.header+Page.rows+Page.footer+FOOTER)
}
