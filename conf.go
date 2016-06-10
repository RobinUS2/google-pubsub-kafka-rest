package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Conf struct {
	JwtJsonKeyPath string
	ProjectId      string
	ListenHost     string
	ListenPort     int
}

func newConf(path string) *Conf {
	c := &Conf{
		ListenPort: 8082, // default port, is the same as confluent kafka rest
	}
	configBytes, configErr := ioutil.ReadFile(path)
	if configErr != nil {
		log.Fatal(fmt.Sprintf("Failed to read configuration: %s", configErr))
	}
	err := json.Unmarshal(configBytes, c)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to read configuration: %s", err))
	}
	return c
}
