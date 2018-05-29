package main

import (
	"github.com/labstack/echo"
	"net/http"
	"log"
)

//ScriptRunner get locations
func ScriptRunner(c echo.Context) error {

	err, out, errout := RunScript("ls -ltr")
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return c.HTML(http.StatusOK, HEADER+out+" ERROUT: "+errout+FOOTER)
}
