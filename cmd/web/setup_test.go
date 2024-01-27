package main

import (
	"net/http"
	"os"
	"testing"
)

type httpHandler struct{}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
