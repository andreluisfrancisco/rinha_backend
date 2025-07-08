FROM golang:1.21 as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app

FROM scratch
WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 9999
CMD ["./app"]