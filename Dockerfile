FROM ubuntu:20.04

RUN apt update && apt install -yq wget && wget https://golang.org/dl/go1.18.linux-amd64.tar.gz

RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

COPY /tests /home

CMD cd /home && go test -v