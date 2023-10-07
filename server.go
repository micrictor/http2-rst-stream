package main

import (
	"fmt"
	"log"
	"net/http"
)

type internalServer struct {
	innerServer *http.Server
	errors      chan error
}

type internalHandler struct {
	response string
}

func (h *internalHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(h.response))
}

func initServer(response string, port int16) *internalServer {
	innerServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: &internalHandler{response},
	}
	server := &internalServer{
		innerServer,
		make(chan error),
	}
	server.Init()
	return server
}

func (i *internalServer) Init() {
	log.Printf("Starting server on %s", i.innerServer.Addr)
	go func() {
		i.errors <- i.innerServer.ListenAndServeTLS(*certFile, *keyFile)
	}()
}
