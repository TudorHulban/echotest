package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.POST("/api/decisions", handlerDecisions)
	e.GET("/api/decisions", handlerDecisionsInDB)
}

// handlerDecisions ...
func handlerDecisions(c echo.Context) error {
	return c.String(http.StatusOK, "Aloha, World!")
}

// handlerDecisionsInDB ...
func handlerDecisionsInDB(c echo.Context) error {
	return c.String(http.StatusOK, "Aloha, World!")
}
