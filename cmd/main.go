package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"path/filepath"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	simulator "github.com/massalabs/thyra-playground-plugin/pkg"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	path, _ := os.Getwd()

	// 	curl -X POST http://localhost:8080/simulate \
	//   -F "files=@./simulator_config.json" \
	//   -F "files=@./main.wasm" \
	//   -H "Content-Type: multipart/form-data"
	router.POST("/simulate", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["files"]

		simulatorPath := filepath.Join(path, "simulator")
		if len(files) != 2 {
			c.String(http.StatusUnauthorized, "/simulate expects 2 files.")
			return
		}
		for _, file := range files {
			log.Println("received file: " + file.Filename)

			if file.Filename != "simulator_config.json" && file.Filename != "main.wasm" {
				log.Fatal("unauthorized filename " + file.Filename)
				c.String(http.StatusUnauthorized, "Unauthorized file names. Should bu simulator_config.json and main.wasm")
				return
			}

			c.SaveUploadedFile(file, filepath.Join(simulatorPath, file.Filename))
		}

		err := simulator.RunSimulation()
		if err != nil {
			log.Fatal(err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		type Res struct {
			LedgerUrl string
			TracesUrl string
		}
		res := &Res{LedgerUrl: c.Request.Host + "/simulator/ledger.json", TracesUrl: c.Request.Host + "/simulator/trace.json"}
		c.PureJSON(http.StatusOK, res)
	})

	router.StaticFile("/simulator/ledger.json", filepath.Join(path, "ledger.json"))
	router.StaticFile("/simulator/trace.json", filepath.Join(path, "trace.json"))

	return router
}

//go:embed static/*
var statics embed.FS

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	if err != nil {
		return false
	}
	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}

func embedStatics(router *gin.Engine) {
	router.Use(static.Serve("/", EmbedFolder(statics, "static")))
}

func main() {
	port := flag.Int("port", 8080, "set listening port")
	flag.Parse()

	router := setupRouter()
	embedStatics(router)

	router.Run("127.0.0.1:" + strconv.Itoa(*port))
}
