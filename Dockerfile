FROM golang:latest AS builder
WORKDIR /build
COPY app .
RUN go get main
RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-s -w -extldflags "-static"' -o procheck


FROM alpine:latest
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=builder /build/procheck .
# CMD bash -c "./procheck -update 5 -port 8080 test1 test2 test3"
