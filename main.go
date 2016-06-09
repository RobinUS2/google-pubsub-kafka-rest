package main

import (
	"io/ioutil"
)

func main() {
	key, _ := ioutil.ReadFile("/tmp/key.json")
	pubsub := newPubSubHelper("test-project", key)
	server := newServer(pubsub)
	server.Start()
}
