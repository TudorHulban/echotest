package main

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
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
	e.POST(url, HandlerPostDecisions)

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
	e.GET(url, HandlerGetDecisions)

	// TODO: add checking of response is JSON array
	apitest.New().
		Handler(e).
		Method(http.MethodGet).
		URL(url).
		Expect(t).
		Status(http.StatusOK).
		End()
}
