FROM golang:latest AS builder
WORKDIR /build
COPY app .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o procheck


FROM alpine:latest
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=builder /build/procheck .
CMD sh -c "./procheck 1 8080 test_metric"
