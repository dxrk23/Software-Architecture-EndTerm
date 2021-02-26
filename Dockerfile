FROM golang:alpine as build-env
ADD . /app
WORKDIR /app

EXPOSE 8080

RUN go run ./cmd/server/avgServer/main.go
RUN go run ./cmd/server/avgServer/main.go

RUN go run ./cmd/client/computeAverage/main.go
RUN go run ./cmd/client/primeDecomposition/main.go

ENTRYPOINT /app
