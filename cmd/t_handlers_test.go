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
		reqBody        string
		statusCodeHTTP int
		respBody       string
	}{
		{testName: "bad, amount as string", reqBody: `{"name": "x", "amount":"100"}`, statusCodeHTTP: http.StatusBadRequest, respBody: ""},
		{testName: "should pass", reqBody: `{"name": "x", "amount":100}`, statusCodeHTTP: http.StatusOK, respBody: `{"decision": true}`},
	}

	e := echo.New()
	addRoutes(e)

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			apitest.New().
				Handler(e).
				Method(http.MethodPost).
				JSON(tc.reqBody).
				URL(url).
				Expect(t).
				Status(tc.statusCodeHTTP).
				Body(tc.respBody).
				End()
		})
	}
}

func TestHandlerGet(t *testing.T) {

}
