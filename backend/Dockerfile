FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bot ./cmd/bot
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bot_manage ./cmd/bot/manage

#####

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
RUN adduser -D -s /bin/sh appuser

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/bot .
COPY --from=builder /app/bot_manage .

# Change ownership to non-root user
RUN chown -R appuser:appuser /root
USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

ENTRYPOINT [ "./server" ]
