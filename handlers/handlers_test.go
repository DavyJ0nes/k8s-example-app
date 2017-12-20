package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := Router()
	ts := httptest.NewServer(r)
	defer ts.Close()

	// TESTING /home
	res, err := http.Get(ts.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /home is wrong. Have %d, want %d.", res.StatusCode, http.StatusOK)
	}

	// TESTING /non-existant
	res, err = http.Get(ts.URL + "/non-existant")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code for /non-existant is wrong. Have %d, want %d.", res.StatusCode, http.StatusNotFound)
	}
}
