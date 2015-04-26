// Package server provides an interface for creating a simple server
package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	host    string
	port    string
	handler http.Handler
	server  *http.Server
}

func NewServer(host string, port string, handler http.Handler) *Server {
	server := &Server{
		host:    host,
		port:    port,
		handler: handler,
	}
	return server
}

func (this *Server) SetHost(host string) {
	this.host = host
}

func (this *Server) SetPort(port string) {
	this.port = port
}

func (this *Server) Address() string {
	return fmt.Sprintf("%s:%s", this.host, this.port)
}

func (this *Server) Start() error {
	this.server = &http.Server{
		Addr:    this.Address(),
		Handler: this.handler,
	}

	return this.server.ListenAndServe()
}
