package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.Lshortfile)

	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":3443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

// package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo"
// 	"github.com/labstack/echo/middleware"
// 	"github.com/peterducai/aidb_portal/cmd/aidbportal"
// 	"golang.org/x/crypto/acme/autocert"
// )

// //EmptyIndex return index.html
// func EmptyIndex(c echo.Context) error {
// 	return c.HTML(http.StatusOK, aidbportal.HEADER+"empty index file"+aidbportal.FOOTER)
// }

// func main() {

// 	e := echo.New()
// 	e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("localhost")
// 	// Cache certificates
// 	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
// 	e.Pre(middleware.HTTPSRedirect()) // redirect to HTTPS
// 	e.Pre(middleware.WWWRedirect())
// 	e.Use(middleware.Recover())
// 	e.Use(middleware.Logger())

// 	//InitDB("postgres://docker:post123@localhost/aidb?sslmode=disable")

// 	//ROUTERS
// 	e.Static("/", "../web/static/")

// 	e.GET("/", EmptyIndex)
// 	e.GET("/locations", aidbportal.GetLocations)
// 	e.GET("/rooms", aidbportal.GetRooms)
// 	e.GET("/racks", aidbportal.GetRacks)

// 	//JUMPS
// 	e.GET("/jumps", aidbportal.ListJumps)
// 	e.POST("/jumps", aidbportal.CreateJump)

// 	//script executor
// 	e.GET("/script/:cmd", aidbportal.ScriptRunner)

// 	e.Logger.Fatal(e.StartTLS(":443", "server.crt", "server.key"))
// 	//e.Logger.Fatal(e.Start(":80"))
// }
