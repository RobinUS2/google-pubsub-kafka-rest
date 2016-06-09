package main

import (
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
func (ps *PubSubHelper) Publish(topic string, data []byte) (bool, error) {
	_, err := pubsub.Publish(ps.ctx, topic, &pubsub.Message{
		Data: data,
	})
	if err == nil {
		return true, nil
	}
	return false, err
}

// Constructor
func newPubSubHelper(projectId string, jwtJsonConfig []byte) *PubSubHelper {
	ps := &PubSubHelper{
		projectId:     projectId,
		jwtJsonConfig: jwtJsonConfig,
	}
	ps.init()
	return ps
}
