package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func initServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routeHandler)

	listenAddr := os.Getenv("HOST_ADDR")
	if listenAddr == "" {
		listenAddr = "127.0.0.1"
	}

	server := http.Server{
		Addr:    fmt.Sprintf("%s:8080", listenAddr),
		Handler: mux,
	}

	go func() {
		server.ListenAndServe()
	}()

	return &server
}

func routeHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "ok")
}
