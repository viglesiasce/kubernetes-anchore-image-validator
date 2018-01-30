FROM golang:1-alpine

ENV NAME=anchore-image-admission-server

WORKDIR /go/src/github.com/viglesiasce/$NAME
COPY . .
RUN cd cmd/$NAME && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /$NAME
CMD ["/$NAME"]
