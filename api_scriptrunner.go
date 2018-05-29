package main

import (
	"github.com/labstack/echo"
	"net/http"
	"log"

)

//ScriptRunner get locations
func ScriptRunner(c echo.Context) error {

	cmd := c.Param("cmd")

	//fmt.Println("cmd:"+cmd)

	err, out, errout := RunScript(cmd)
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return c.HTML(http.StatusOK, HEADER+out+" ERROUT: "+errout+FOOTER)
}
