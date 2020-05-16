FROM golang:1.13 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ADD . /usr/src/app
WORKDIR /usr/src/app
RUN go build -o research main.go

# runtime image
FROM alpine

RUN mkdir -p /data/pdf/handout
COPY --from=buildImage /usr/src/app/research /
ENTRYPOINT ["/research"]
