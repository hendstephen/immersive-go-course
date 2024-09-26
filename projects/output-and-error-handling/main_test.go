package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func captureOutput(f func()) (string, error) {
	origOut := os.Stdout
	origErr := os.Stderr

	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()

	os.Stdout = wOut
	os.Stderr = wErr

	f()

	os.Stdout = origOut
	os.Stderr = origErr

	wOut.Close()
	wErr.Close()

	out, _ := io.ReadAll(rOut)
	err, _ := io.ReadAll(rErr)

	return string(out), errors.New(string(err))
}

func TestCallServer_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Today it will be sunny!"))
	}))
	defer server.Close()

	out, _ := captureOutput(func() {
		callServer(server.URL)
	})

	fmt.Printf("Output: %s\n", out)
}
