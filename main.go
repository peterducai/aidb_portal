package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"

	"github.com/peterducai/aidb_portal/cmd/aidbportal"
)

//EmptyIndex return index.html
func EmptyIndex(c echo.Context) error {
	return c.HTML(http.StatusOK, HEADER+"empty index file"+FOOTER)
}

func main() {

	e := echo.New()
	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<DOMAIN>")
	// Cache certificates
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Pre(middleware.HTTPSRedirect()) // redirect to HTTPS
	e.Pre(middleware.WWWRedirect())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	//InitDB("postgres://docker:post123@localhost/aidb?sslmode=disable")

	//ROUTERS
	e.Static("/", "../web/static/")

	e.GET("/", EmptyIndex)
	e.GET("/locations", aidbportal.GetLocations)
	e.GET("/rooms", GetRooms)
	e.GET("/racks", GetRacks)

	//JUMPS
	e.GET("/jumps", ListJumps)
	e.POST("/jumps", CreateJump)

	//script executor
	e.GET("/script/:cmd", ScriptRunner)

	e.Logger.Fatal(e.StartAutoTLS(":443"))
	//e.Logger.Fatal(e.Start(":80"))
}
