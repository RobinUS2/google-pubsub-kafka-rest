package main

import (
	"io/ioutil"
)

func main() {
	key, _ := ioutil.ReadFile("/tmp/key.json")
	pubsub := newPubSubHelper("robin-1225", key)
	server := newServer(pubsub)
	server.Start()
}
