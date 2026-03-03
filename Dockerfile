FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY grape-shared ./grape-shared
COPY pkg ./pkg
COPY api ./api
COPY services/auth-service ./services/auth-service
WORKDIR /app/services/auth-service
RUN go build -o /auth-service ./cmd/auth-service/main.go

FROM alpine:3.19 AS runner
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /auth-service .
COPY --from=builder /app/services/auth-service/migrations ./migrations
EXPOSE 8060
ENTRYPOINT ["./auth-service"]
