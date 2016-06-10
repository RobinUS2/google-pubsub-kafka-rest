package main

import (
	"flag"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "conf", "/etc/pubsub_kafa_rest.json", "Path to configuration JSON (not the key)")
	flag.Parse()
}

func main() {
	// Config
	conf := newConf(configPath)

	// Pubsub
	pubsub := newPubSubHelper(conf)

	// server
	server := newServer(conf, pubsub)
	server.Start()
}
