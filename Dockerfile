# ---- build ----
FROM golang:1.22-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# ---- run ----
FROM alpine:3.20
WORKDIR /app

RUN adduser -D appuser
USER appuser

COPY --from=build /app/server /app/server

EXPOSE 8080
ENTRYPOINT ["/app/server"]