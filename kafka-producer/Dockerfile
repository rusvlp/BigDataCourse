
FROM golang:1.21.4-bookworm

ADD /kafka-producer-application /kpa
ADD /start.sh /start.sh

WORKDIR /kpa/internal

RUN go build main.go
RUN mv main /

WORKDIR /



CMD ls && pwd && /start.sh 