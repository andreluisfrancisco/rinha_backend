FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o rinha-gateway

FROM alpine:3.18

RUN apk add --no-cache ca-certificates wget

RUN addgroup -g 10001 -S nonroot && \
    adduser -u 10001 -S -G nonroot -h /home/nonroot nonroot

WORKDIR /app

COPY --from=builder /app/rinha-gateway .

RUN chmod +x rinha-gateway

USER nonroot:nonroot

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

ENTRYPOINT ["./rinha-gateway"]