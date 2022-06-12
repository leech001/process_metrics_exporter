FROM golang:latest AS builder

WORKDIR /build
COPY app .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o

FROM golang:latest
WORKDIR /app
COPY --from=builder /build/main .
CMD sh -c "./main"
