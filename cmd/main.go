package main

import (
	"context"
	"net/http"

	"github.com/TudorHulban/echotest/pkg/logic"
	"github.com/TudorHulban/echotest/pkg/models"
	"github.com/TudorHulban/echotest/pkg/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	url = "/api/decisions"
)

// RequestDecision Used to bind JSON body in handler.
type RequestDecision struct {
	Name   string
	Amount int
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.HideBanner = true

	addRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func addRoutes(e *echo.Echo) {
	e.POST(url, handlerDecisions)
	e.GET(url, handlerDecisionsInDB) // requirement not in line with CRUD
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

// handlerDecisionsInDB Saves decision to database.
func handlerDecisionsInDB(c echo.Context) error {
	model := new(RequestDecision)

	if errBind := c.Bind(model); errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": errBind.Error()})
	}

	// TODO: add request ID
	if errInsert := repository.GetInstance().Create(context.Background(), &models.Decision{Name: model.Name, Amount: model.Amount}); errInsert != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": errInsert.Error()})
	}

	return c.String(http.StatusOK, "")
}
