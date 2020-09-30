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
		{testName: "decision true", reqBody: `{"name": "x", "amount":100}`, statusCodeHTTP: http.StatusOK, respBody: `{"decision": true}`},
		{testName: "decision false", reqBody: `{"name": "x", "amount":10001}`, statusCodeHTTP: http.StatusOK, respBody: `{"decision": false}`},
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
	e := echo.New()
	addRoutes(e)

	apitest.New().
		Handler(e).
		Method(http.MethodGet).
		JSON(`{"requestid":"100","name": "x", "amount":100, "answer": true}`).
		URL(url).
		Expect(t).
		Status(http.StatusOK).
		Body("").
		End()
}
