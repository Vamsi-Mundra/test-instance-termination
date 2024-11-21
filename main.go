package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	router.POST("/data", func(c *gin.Context) {
		// Read the incoming text data
		body, err := c.GetRawData()
		if err != nil {
			// Log an error if reading the body fails
			c.String(http.StatusBadRequest, "Failed to read body")
			return
		}

		// Log the received text data
		fmt.Printf("Received data: %s\n", string(body))

		// Return an OK response
		c.String(http.StatusOK, "OK")
	})

	router.Run(":" + port)
}
