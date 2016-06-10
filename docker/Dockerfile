FROM centos:7
MAINTAINER "Robin Verlangen"
RUN yum install -y git golang
RUN git clone https://github.com/RobinUS2/google-pubsub-kafka-rest.git
RUN ./google-pubsub-kafka-rest/build.sh
CMD ["./google-pubsub-kafka-rest/google-pubsub-kafka-rest"]
