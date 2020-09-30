package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Request struct {
	Name   string
	Amount string
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HideBanner = true

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
