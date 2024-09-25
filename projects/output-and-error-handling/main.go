package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Client struct {
	BaseUrl string
	Out     io.Writer
	Errout  io.Writer
}

func callServer(url string) {
	fmt.Fprintf(os.Stdout, "Making request to %s...\n", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to make request: %v\n", err)
		os.Exit(1)
	} else if res.StatusCode == 429 {
		retryAfter := res.Header.Get("Retry-After")
		if retryAfter == "" {
			fmt.Fprintf(os.Stderr, "Server is too busy, but no Retry-After header was provided\n")
		} else {
			duration := time.Second
			if parseTimestamp, err := time.Parse(http.TimeFormat, retryAfter); err == nil {
				duration = time.Until(parseTimestamp)
			} else if parseDelay, err := strconv.Atoi(retryAfter); err == nil {
				duration = time.Duration(parseDelay) * time.Second
			}

			if duration > time.Second {
				fmt.Fprintf(os.Stdout, "Server is too busy, retrying after %s\n", duration)
			}

			time.Sleep(duration)
			callServer(url)
		}
	} else {
		fmt.Fprintf(os.Stdout, "Response: ")
		io.Copy(os.Stdout, res.Body)
		fmt.Fprintln(os.Stdout)
	}
}

func main() {
	callServer("http://localhost:8080/")
}
