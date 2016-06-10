#!/bin/bash
docker run --name pubsubrest \
	-v /etc/pubsub_kafa_rest.json:/etc/pubsub_kafa_rest.json \
	-v /tmp/key.json:/tmp/key.json \
	-d \
	--restart=always \
	--ulimit nofile=262144:262144 \
	--memory="1G" \
	-p 8083:8083 \
	pubsubrest
