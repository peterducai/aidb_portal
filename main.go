package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type server struct {
	db *sql.DB
}

func main() {

	fmt.Println("AIDB Portal server")
	fmt.Println("version 0.1")

	// DB PART
	db, err := sql.Open("postgres", "database=aidb user=docker password=post123 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//s := server{db: db}

	// SERVER PART
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetIndex)
	mux.HandleFunc("/locations", GetLocations)

	// TLS config
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	// SERVER config
	srv := &http.Server{
		Addr:         ":443",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}
