package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	simulator "github.com/massalabs/thyra-playground-plugin/pkg"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	router.Static("/", "./static")

	router.POST("/simulate", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			if file.Filename != "simulator_config.json" && file.Filename != "main.wasm" {
				log.Fatal("unauthorized filename " + file.Filename)
			}
			c.SaveUploadedFile(file, "./pkg/"+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

		simulator.RunSimulation()

	})

	return router
}

func main() {
	port := flag.Int("port", 8080, "set listening port")
	flag.Parse()

	router := setupRouter()

	simulator.RunSimulation()

	router.Run("127.0.0.1:" + strconv.Itoa(*port))
}
