package main

import (
	"net/http"
)

//GetIndex get locations
func GetIndex(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is an AIDB Portal server.\nSee <a href=\"locations\">locations</a>\n"))
}
