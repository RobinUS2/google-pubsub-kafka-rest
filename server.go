package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
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
			fooHandler(s, ctx)
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
func fooHandler(s *Server, ctx *fasthttp.RequestCtx) {
	res, err := s.pubsub.Publish("testpubsubkafkarest", []byte("Hello"))
	log.Printf("Res %v err %v", res, err)
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
