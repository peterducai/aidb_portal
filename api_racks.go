package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//GetRacks get locations
func GetRacks(c echo.Context) error {

	racks, err := AllRacks()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "db connection error")
	}
	for _, rm := range racks {
		fmt.Println(rm.id, rm.rackName, rm.rackSize, rm.purchaseOrder, rm.rack)
	}

	return c.String(http.StatusOK, "racks!")
}
