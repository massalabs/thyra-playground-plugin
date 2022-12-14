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

func setupRouter(pluginPath string) *gin.Engine {

	router := gin.Default()

	// 	curl -X POST http://localhost:8080/simulate \
	//   -F "files=@./simulator_config.json" \
	//   -F "files=@./main.wasm" \
	//   -H "Content-Type: multipart/form-data"
	router.POST("/simulate", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			log.Fatal(err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		files := form.File["files"]

		simulatorPath := filepath.Join(pluginPath, "simulator")
		if len(files) != 2 {
			msg := "/simulate expects 2 files."
			log.Fatal(msg)
			c.String(http.StatusBadRequest, msg)
			return
		}
		for _, file := range files {
			log.Println("received file: " + file.Filename)

			if file.Filename != "simulator_config.json" && file.Filename != "main.wasm" {
				log.Fatal("unauthorized filename " + file.Filename)
				c.String(http.StatusBadRequest, "Unauthorized file names. Should bu simulator_config.json and main.wasm")
				return
			}

			c.SaveUploadedFile(file, filepath.Join(simulatorPath, file.Filename))
		}

		err = simulator.RunSimulation(pluginPath)
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

	ledgerPath := filepath.Join(pluginPath, "ledger.json")
	tracePath := filepath.Join(pluginPath, "trace.json")

	// We create empty files because the frontend use these files presence to detect if the simulator is available...
	// These files will be written later by the simulator
	ledgerFile, _ := os.Create(ledgerPath)
	ledgerFile.WriteString("{}")
	defer ledgerFile.Close()
	traceFile, _ := os.Create(tracePath)
	traceFile.WriteString("{}")
	defer traceFile.Close()

	router.StaticFile("/simulator/ledger.json", ledgerPath)
	router.StaticFile("/simulator/trace.json", tracePath)

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
	path, _ := os.Getwd()
	pluginPath := flag.String("path", path, "set listening port")
	flag.Parse()

	router := setupRouter(*pluginPath)
	embedStatics(router)

	router.Run("127.0.0.1:" + strconv.Itoa(*port))
}
