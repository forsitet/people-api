FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/app

FROM alpine:3.21.3

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/server .

COPY --from=builder /app/internal /root/internal

COPY --from=builder /app/shared /root/shared

COPY --from=builder /app/swagger.html ./swagger.html

COPY --from=builder /app/shared/api/bundles/people.openapi.v1.bundle.yaml ./people.openapi.v1.bundle.yaml

COPY --from=builder /app/internal/repository/migrations ./migrations

COPY --from=builder /app/.env ./.env

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["./server"] 
