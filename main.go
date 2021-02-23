package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.StaticFS("/", static)
	log.Println("starting server on: http://localhost" + port)
	router.Run(port)

	fmt.Println("Done ...")
}
