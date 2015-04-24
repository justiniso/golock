// Package server provides an interface for creating a simple server
package server

import (
    "net/http"
)

type Server struct {
    var server = &http.Server
    var address string
    var port string
}

func (this *Server) SetAddress(address string) {
    this.address = address
}

func (this *Server) SetPort(port string) {
    this.port = port
}

func (this *Server) Start() error {
    this.server = &http.Server{
        Addr: this.address,
        Handler: this,
    }

    return this.server.ListenAndServe()
}