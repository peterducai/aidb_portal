package main

import (
	"net/http"
)

//GetLocations get locations
func GetLocations(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Locations.\n"))
}
