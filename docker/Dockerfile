FROM centos:7
MAINTAINER "Robin Verlangen"
RUN yum install -y git wget
RUN wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz && tar -C /usr/local -zxvf go*.tar.gz && cp /usr/local/go/bin/go /usr/bin/go
PWD /usr/local
RUN git clone https://github.com/RobinUS2/google-pubsub-kafka-rest.git
RUN ./google-pubsub-kafka-rest/build.sh
CMD ["/usr/local/google-pubsub-kafka-rest/google-pubsub-kafka-rest"]
