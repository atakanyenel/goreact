package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

const port = ":8090"

//go:generate npm run --prefix frontend build
//go:embed frontend/build
var content embed.FS

func main() {
	fmt.Println("Go React demo app, v1")

	dev := flag.Bool("dev", false, "isdev")
	flag.Parse()

	static := http.FileServer(http.Dir("./frontend/build"))
	if !*dev {
		fsys, err := fs.Sub(content, "frontend/build")
		if err != nil {
			panic(err)
		}
		static = http.FileServer(http.FS(fsys))
	}
	http.Handle("/", static)
	log.Println("starting server on: http://localhost" + port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

	fmt.Println("Done ...")
}
