package main

// see documentation at https://github.com/gavv/httpexpect/tree/master

import(
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gavv/httpexpect/v2"
	"github.com/timoruohomaki/open311togo/server"
)

func TestOpen311GetTime(t *testing.T) {
	testHandler := server.Handlers.GetServices()

	server := httptest.NewServer(testHandler)

	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	e.GET("/open311/rest/v1/time").
		Expect().
		Status(http.StatusOK).JSON().Array().IsEmpty()

}