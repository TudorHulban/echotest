package main

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/steinfletcher/apitest"
)

func TestHandlerPost(t *testing.T) {
	tt := []struct {
		testName       string
		jsonBody       string
		statusCodeHTTP int
	}{
		{testName: "bad, amount as string", jsonBody: `{"name": "x", "amount":"100"}`, statusCodeHTTP: http.StatusBadRequest},
		{testName: "should pass", jsonBody: `{"name": "x", "amount":100}`, statusCodeHTTP: http.StatusOK},
	}

	e := echo.New()
	addRoutes(e)

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			apitest.New().
				Handler(e).
				Method(http.MethodPost).
				JSON(tc.jsonBody).
				URL(url).
				Expect(t).
				Status(tc.statusCodeHTTP).
				End()
		})
	}
}

func TestHandlerGet(t *testing.T) {

}
