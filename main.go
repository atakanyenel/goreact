package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
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
	var static http.FileSystem = http.Dir("./frontend/build")
	if !*dev { // use embedded in prod
		fsys, err := fs.Sub(content, "frontend/build")
		if err != nil {
			panic(err)
		}
		static = http.FS(fsys)
	}

	// Serve static files
	http.Handle("/", http.FileServer(static))
	http.ListenAndServe(port, nil)

	fmt.Println("Done ...")
}
