package aidbportal

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"text/template"
)

//MyData my data
type MyData struct {
	mytool string
}

//Index get index.html
func Index(w http.ResponseWriter, req *http.Request) {
	// all calls to unknown url paths should return 404
	if req.URL.Path != "/" {
		log.Printf("404: %s", req.URL.String())
		http.NotFound(w, req)
		return
	}
	http.ServeFile(w, req, "templates/index.html")
}

//GetIndex get locations
func GetIndex(w http.ResponseWriter, req *http.Request) {

	md := MyData{mytool: "index.html"}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	//w.Write([]byte("This is an AIDB Portal server.\nSee <a href=\"locations\">locations</a>\n"))
	lp := filepath.Join(dir+"/templates/", "index.html")
	fp := filepath.Join(dir+"/templates/", filepath.Clean(req.URL.Path))

	fmt.Println(lp)
	fmt.Println(fp)
	//fmt.Println()

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, req)
			fmt.Println("templates NOT found")
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, req)
		fmt.Println("request is for DIR! ..404")
		return
	}

	t, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := t.ExecuteTemplate(w, md.mytool, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
