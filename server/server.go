package server

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var addr string
var shutdown bool
var listener net.Listener

// Start() the web server
func Start() {
	addr = "127.0.0.1:8000"
	server := &http.Server{
		Addr:         addr,
		Handler:      nil,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	// We don't use ListenAndServe because it lacks a way to close the listener
	fmt.Printf("HTTP listening on TCP4 %v", addr)
	var err error
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("HTTP failed to start TCP4 listener: %v", err)
		// TODO More graceful early-shutdown procedure
		panic(err)
	}

	err = server.Serve(listener)
	if shutdown {
		fmt.Printf("HTTP server shutting down on request")
	} else if err != nil {
		fmt.Printf("HTTP server failed: %v", err)
	}
}

func Stop() {
	fmt.Printf("HTTP shutdown requested")
	shutdown = true
	if listener != nil {
		listener.Close()
	} else {
		fmt.Printf("HTTP listener was nil during shutdown")
	}
}
