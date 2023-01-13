package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
