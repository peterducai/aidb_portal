package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//EmptyIndex return index.html
func EmptyIndex(c echo.Context) error {
	return c.HTML(http.StatusOK, Header+Content+Footer)
}

func main() {

	e := echo.New()
	// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<DOMAIN>")
	// Cache certificates
	//e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	//e.Pre(middleware.HTTPSRedirect())      // redirect to HTTPS
	e.Pre(middleware.WWWRedirect())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	InitDB("postgres://docker:post123@localhost/aidb?sslmode=disable")

	//ROUTERS
	e.Static("/", "../web/static/")

	e.GET("/", EmptyIndex)
	e.GET("/locations", GetLocations)
	e.GET("/rooms", GetRooms)

	//e.Logger.Fatal(e.StartAutoTLS(":443"))
	e.Logger.Fatal(e.Start(":80"))
}
