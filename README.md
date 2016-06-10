# google-pubsub-kafka-rest
Implements the basic features in Confluent kafka rest but produces to Google PubSub

Why would I use this?
-------------
It is developed as an easy way to migrate for Kafka with the [Confluent Kafka REST API](https://github.com/confluentinc/kafka-rest) to [Google PubSub](https://cloud.google.com/pubsub/overview).

Configuration
-------------
Place a configuration file in `/etc/pubsub_kafa_rest.json`

```
{
	"JwtJsonKeyPath": "/tmp/key.json",
	"ProjectId": "your-project-123"
}
```

Place the Google IAM obtained service account JSON key in `/tmp/key.json` (or the path you specified in the config)

Examples
------------
To send the string value "Hello" into the topic "test1" with cURL on localhost with default configs:
```
curl -XPOST --data '{"records":[{"value":"SGVsbG8="}]}' http://localhost:8082/topics/test1
```
