FROM golang:1.13 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ADD . /usr/src/app
WORKDIR /usr/src/app
RUN go build -o grpc-gateway main.go

# runtime image
FROM alpine

COPY --from=buildImage /usr/src/app/grpc-gateway /
ENTRYPOINT ["/grpc-gateway"]
