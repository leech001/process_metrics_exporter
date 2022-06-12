FROM golang:latest AS builder
WORKDIR /build
COPY app .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o procheck


FROM alpine:latest
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=builder /build/procheck .
# CMD sh -c "./procheck -update 5 -port 8080 test1 test2 test3"
