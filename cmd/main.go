package main

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TudorHulban/echotest/pkg/logic"
	"github.com/TudorHulban/echotest/pkg/models"
	"github.com/TudorHulban/echotest/pkg/repository"
	guuid "github.com/google/uuid"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	url = "/api/decisions"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return customGenerator()
		},
	}))

	// commented for easy testing
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	e.HideBanner = true

	e.POST(url, HandlerPostDecisions)
	e.GET(url, HandlerGetDecisions)
	e.Logger.Fatal(e.Start(":1323"))
}

func customGenerator() string {
	return guuid.New().String()
}

// HandlerPostDecisions Serving POST requests.
//
// Manual test:
// curl -X POST http://localhost:1323/api/decisions  -H 'Content-Type: application/json' -d '{"name":"X","amount":100}' -H 'Authorization: Basic am9lOnNlY3JldA=='
func HandlerPostDecisions(c echo.Context) error {
	log.Println("Request ID:", c.Request().Header.Get(echo.HeaderXRequestID))

	model := new(models.Decision)

	// TODO: add request validation
	if errBind := c.Bind(model); errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": errBind.Error()})
	}

	decision, errDecision := logic.DecisionAmount(model.Amount)
	if errDecision != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": errDecision.Error()})
	}

	model.RequestID = c.Response().Writer.Header().Get(echo.HeaderXRequestID)
	model.Answer = decision

	if errInsert := repository.GetInstance().Create(context.Background(), model); errInsert != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": errInsert.Error()})
	}

	return c.JSON(http.StatusOK, map[string]bool{"decision": decision})
}

// handlerDecisionsInDB Saves decision to database.
//
// Manual test:
// curl -X GET http://localhost:1323/api/decisions  -H 'Content-Type: application/json' -H 'Authorization: Basic am9lOnNlY3JldA=='
func HandlerGetDecisions(c echo.Context) error {
	records, errFind := repository.GetInstance().FindAll(context.Background())
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": errFind.Error()})
	}

	log.Println("retrieved records:", records)

	decisions := make([]*models.Decision, len(*records))
	for ix, v := range *records {
		decisions[ix] = &v
	}

	log.Println("massaged records:", decisions)

	data, errMa := json.Marshal(decisions)
	if errMa != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"status": errMa.Error()})
	}

	return c.JSON(http.StatusOK, string(data))
}
