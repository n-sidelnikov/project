FROM ubuntu:20.04

RUN export DEBIAN_FRONTEND=noninteractive && apt update && apt install -yq golang

COPY /tests/hello_test.go /

CMD go test -run hello_test.go