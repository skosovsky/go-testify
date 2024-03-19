package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	statusExpected := http.StatusOK

	answer := responseRecorder.Body

	require.Equal(t, status, statusExpected)
	assert.NotEmpty(t, answer)
}

func TestMainHandlerWhenMissingCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	statusExpected := http.StatusBadRequest

	answer := responseRecorder.Body.String()
	answerContains := "count missing"

	require.Equal(t, status, statusExpected)
	assert.Contains(t, answer, answerContains)
}

func TestMainHandlerWhenWrongCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	statusExpected := http.StatusBadRequest

	answer := responseRecorder.Body.String()
	answerContains := "wrong count value"

	require.Equal(t, status, statusExpected)
	assert.Contains(t, answer, answerContains)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=paris", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	statusExpected := http.StatusBadRequest

	answer := responseRecorder.Body.String()
	answerContains := "wrong city value"

	require.Equal(t, status, statusExpected)
	assert.Contains(t, answer, answerContains)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	statusExpected := http.StatusOK

	answer := responseRecorder.Body.String()
	list := strings.Split(answer, ",")
	count := len(list)
	countExpected := totalCount

	require.Equal(t, status, statusExpected)
	assert.Equal(t, count, countExpected)
}
