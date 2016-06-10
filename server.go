package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type Server struct {
	requestHandler func(ctx *fasthttp.RequestCtx)
	pubsub         *PubSubHelper
	conf           *Conf
	router         *routing.Router
}

// Init
func (s *Server) init() {
	s.router = routing.New()
	s.router.Post("/topics/<topic>", func(c *routing.Context) error {
		// Params
		topic := c.Param("topic")
		body := c.PostBody()

		// Receive data
		var data PostTopic
		jsonErr := json.Unmarshal(body, &data)
		if jsonErr != nil {
			// Failed to unmarshal
			log.Printf("Json invalid %s", jsonErr)
			fmt.Fprintf(c, "{\"error_code\":400,\"message\":\"HTTP 400 Bad Request: invalid json\"}")
			c.SetStatusCode(400)
			return nil
		}

		// Convert data
		byteArr := make([][]byte, len(data.Records))
		for i, record := range data.Records {
			recordBytes, b64e := base64.StdEncoding.DecodeString(record.Value)
			if b64e != nil {
				log.Printf("Failed to base64 decode: %s", b64e)
			}
			byteArr[i] = recordBytes
		}

		// Write to pubsub
		_, err := s.pubsub.Publish(topic, byteArr...)
		if err != nil {
			// Failed to publish
			log.Printf("Unable to produce to PubSub %s", err)
			fmt.Fprintf(c, "{\"error_code\":503,\"message\":\"HTTP 503 Service Unavailable: unable to produce to PubSub\"}")
			c.SetStatusCode(503)
			return nil
		}

		// Static OK json response as we don't need / have this information
		fmt.Fprintf(c, "{\"offsets\":[{\"partition\":0,\"offset\":0,\"error_code\":null,\"error\":null}],\"key_schema_id\":null,\"value_schema_id\":null}")
		return nil
	})
}

// Start
func (s *Server) Start() {
	listenStr := fmt.Sprintf("%s:%d", s.conf.ListenHost, s.conf.ListenPort)
	log.Printf("Starting server at %s", listenStr)
	panic(fasthttp.ListenAndServe(listenStr, s.router.HandleRequest))
}

// Constructor
func newServer(conf *Conf, pubsub *PubSubHelper) *Server {
	s := &Server{
		conf:   conf,
		pubsub: pubsub,
	}
	s.init()
	return s
}
