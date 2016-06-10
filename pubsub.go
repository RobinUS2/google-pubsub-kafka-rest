package main

import (
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/pubsub"
)

type PubSubHelper struct {
	ctx           context.Context
	projectId     string
	jwtJsonConfig []byte
}

// Init
func (ps *PubSubHelper) init() {
	conf, err := google.JWTConfigFromJSON(
		ps.jwtJsonConfig,
		pubsub.ScopeCloudPlatform,
		pubsub.ScopePubSub,
	)
	if err != nil {
		log.Fatal(err)
	}
	ps.ctx = cloud.NewContext(ps.projectId, conf.Client(oauth2.NoContext))
}

// Write
func (ps *PubSubHelper) Publish(topic string, data ...[]byte) (bool, error) {
	msgs := make([]*pubsub.Message, len(data))
	for i, dataElm := range data {
		msgs[i] = &pubsub.Message{
			Data: dataElm,
		}
	}
	_, err := pubsub.Publish(ps.ctx, topic, msgs...)
	if err == nil {
		return true, nil
	}
	return false, err
}

// Constructor
func newPubSubHelper(conf *Conf) *PubSubHelper {
	// Read key
	key, keyErr := ioutil.ReadFile(conf.JwtJsonKeyPath)
	if keyErr != nil {
		log.Fatalf("Failed to read JSON key: %s", keyErr)
	}

	// Construct helper
	ps := &PubSubHelper{
		projectId:     conf.ProjectId,
		jwtJsonConfig: key,
	}

	// Init
	ps.init()
	return ps
}
