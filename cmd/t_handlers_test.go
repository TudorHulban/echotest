package main

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

func TestHandlerPost(t *testing.T) {
	tt := []struct {
		testName       string
		httpMethod     string
		reqURL         string
		statusCodeHTTP int
	}{
		{testName: "No Credentials", httpMethod: http.MethodPost, reqURL: "", statusCodeHTTP: http.StatusBadRequest},
	}

	s := echo.New()

	if assert.Nil(t, errCo) {
		for _, tc := range tt {
			t.Run(tc.testName, func(t *testing.T) {
				apitest.New().
					Handler(s.engine).
					Method(tc.httpMethod).
					URL(tc.reqURL).
					FormData("usercode", tc.usercode).
					FormData("password", tc.password).
					Expect(t).
					Status(tc.statusCodeHTTP).
					End()
			})
		}
	}
}

func TestHandlerGet(t *testing.T) {

}
