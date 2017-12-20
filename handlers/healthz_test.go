package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	// TESTING /healthz Status Code
	w := httptest.NewRecorder()
	healthz(w, nil)
	resp := w.Result()

	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d", have, want)
	}
}
