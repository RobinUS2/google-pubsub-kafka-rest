#!/bin/bash
docker run --name pubsubrest \
	-v /etc/pubsub_kafa_rest.json:/etc/pubsub_kafa_rest.json \
	-v /tmp/key.json:/tmp/key.json \
	-d \
	--restart=always \
	--ulimit nofile=262144:262144 \
	--memory="1G" \
	--net=host \ # Probably want to disable this in production
	pubsubrest
