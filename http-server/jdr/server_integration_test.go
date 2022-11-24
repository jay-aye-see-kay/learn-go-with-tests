package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWindsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	respose := httptest.NewRecorder()
	server.ServeHTTP(respose, newGetScoreRequest(player))
	assertStatus(t, respose.Code, http.StatusOK)

	assertResponseBody(t, respose.Body.String(), "3")
}
