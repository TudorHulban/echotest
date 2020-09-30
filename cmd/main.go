package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RequestDecision Used to bind JSON body in handler.
type RequestDecision struct {
	Name   string
	Amount string
}

// ResponseDecision Used in response.
type ResponseDecision struct {
	Decision bool `json:"decision`
}

// CustomValidator Used for validating the received data in handler.
type CustomValidator struct {
	validator *validator.Validate
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HideBanner = true

	e.POST("/api/decisions", handlerDecisions)
	e.GET("/api/decisions", handlerDecisionsInDB)

	e.Logger.Fatal(e.Start(":1323"))
}

// handlerDecisions Serving POST requests.
func handlerDecisions(c echo.Context) error {
	model := new(RequestDecision)

	if errBind = c.Bind(model); errBind != nil {
		return c.Error(errBind)
	}

	return c.String(http.StatusOK, "Aloha, World!")
}

// handlerDecisionsInDB ...
func handlerDecisionsInDB(c echo.Context) error {
	return c.String(http.StatusOK, "Aloha, World!")
}
