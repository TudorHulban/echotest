package main

import (
	"log"
	"net/http"

	"github.com/TudorHulban/echotest/pkg/logic"
	"github.com/TudorHulban/echotest/pkg/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	url         = "/api/decisions"
	mongoURL    = "mongodb://localhost:27017"
	mongoDBName = "decisions"
)

// RequestDecision Used to bind JSON body in handler.
type RequestDecision struct {
	Name   string
	Amount int
}

func main() {
	config := &repository.DBConfig{DatabaseName: mongoDBName, DBUrl: mongoURL}
	helper, errRepo := repository.NewClient(config)
	if errRepo != nil {
		log.Fatal(errRepo.Error())
	}

	errCo := helper.Connect()
	if errCo != nil {
		log.Fatal("Cound not connect to mongo {} ", errCo.Error())
	}

	dbHandler := repository.NewDatabase(config, helper)
	dbHandler.Client().StartSession()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	e.HideBanner = true

	e.POST(url, handlerDecisions)
	e.GET(url, handlerDecisionsInDB)

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
	model := new(RequestDecision)

	if errBind := c.Bind(model); errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": errBind.Error()})
	}

	// TODO: model is populated , can be saved in mongo

	return c.String(http.StatusOK, "Aloha, World!")
}
