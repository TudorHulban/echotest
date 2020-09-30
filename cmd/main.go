package main

import (
	"net/http"

	"github.com/TudorHulban/echotest/pkg/logic"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const URL = "/api/decisions"

// RequestDecision Used to bind JSON body in handler.
type RequestDecision struct {
	Name   string
	Amount int
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

	e.POST(URL, handlerDecisions)
	e.GET(URL, handlerDecisionsInDB)

	e.Logger.Fatal(e.Start(":1323"))
}

// handlerDecisions Serving POST requests.
//
// Manual test:
// curl -X POST http://localhost:1323/api/decisions  -H 'Content-Type: application/json' -d '{"name":"X","amount":100}'
func handlerDecisions(c echo.Context) error {
	model := new(RequestDecision)

	// TODO: add validation

	if errBind := c.Bind(model); errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": errBind.Error()})
	}

	decision, errDecision := logic.DecisionAmount(model.Amount)
	if errDecision != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": errDecision.Error()})
	}

	return c.JSON(http.StatusOK, map[string]bool{"decision": decision})
}

// handlerDecisionsInDB ...
func handlerDecisionsInDB(c echo.Context) error {
	return c.String(http.StatusOK, "Aloha, World!")
}
