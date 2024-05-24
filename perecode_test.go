package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenResponseCorrect(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.NotEmpty(t, responseRecorder.Body.String())
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
}

func TestMainHandlerWhenCityDontCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=zvenigorod", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.NotEmpty(t, http.StatusBadRequest, responseRecorder.Body.String())
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, totalCount)
}
