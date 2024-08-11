# build stage
FROM golang:1.22.5-alpine3.20 AS builder


WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download
COPY internal /app/internal
COPY cmd /app/cmd
RUN go build -o /app/cmd/main /app/cmd/main.go

# final stage
FROM alpine:3.20 AS runtime

WORKDIR /app

COPY --from=builder /app/cmd/main /app/main
COPY assets /app/assets
COPY web /app/web

EXPOSE 8080

CMD ["./main"]
