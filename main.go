package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var port = ":8090"

//go:generate npm run --prefix frontend build
//go:embed frontend/dist
var content embed.FS

func main() {
	fmt.Println("Go React demo app, v1")

	r := mux.NewRouter()
	dev := flag.Bool("dev", false, "isdev")
	flag.Parse()
	if *dev {
		r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))
	} else {
		fsys, err := fs.Sub(content, "frontend/dist")
		if err != nil {
			panic(err)
		}
		r.PathPrefix("/").Handler(http.FileServer(http.FS(fsys)))
	}
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Println("starting server on: http://localhost" + port)
	err := http.ListenAndServe(port, loggedRouter)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

	fmt.Println("Done ...")
}
