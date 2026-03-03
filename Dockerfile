FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o /auth-service ./cmd/auth-service/main.go

FROM alpine:3.19 AS runner
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /auth-service .
COPY --from=builder /app/services/auth-service/config.yml ./config.yml
COPY --from=builder /app/services/auth-service/migrations ./migrations
EXPOSE 8060
ENTRYPOINT ["./auth-service"]
