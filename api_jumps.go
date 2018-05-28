package main

import (
	"net/http"

	"github.com/labstack/echo"
)

//CreateJump get locations
func CreateJump(c echo.Context) error {

	//jmps, err := AllJumps()
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return c.String(http.StatusBadRequest, "issue with DB connection - jumps call")
	//}

	jumpname := c.FormValue("jumpname")


	return c.HTML(http.StatusOK, HEADER+jumpname+FOOTER)
}
