
# Builder
FROM golang:1.25.5 AS builder
WORKDIR /opt
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/api

# Runner
FROM scratch
WORKDIR /opt
COPY --from=builder /opt/server .
CMD ["/opt/server"]
