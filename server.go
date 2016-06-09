package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type Server struct {
	requestHandler func(ctx *fasthttp.RequestCtx)
	pubsub         *PubSubHelper
}

// Init
func (s *Server) init() {
	s.requestHandler = func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/foo":
			fooHandler(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
}

// Start
func (s *Server) Start() {
	fasthttp.ListenAndServe(":8081", s.requestHandler)
}

// Foo
func fooHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

// Constructor
func newServer(pubsub *PubSubHelper) *Server {
	s := &Server{
		pubsub: pubsub,
	}
	s.init()
	return s
}
