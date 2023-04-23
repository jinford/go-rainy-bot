####################### Build stage #######################
FROM golang:1.19 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd/go-rainy-bot
# alpineでGoのバイナリを実行するには CGO_ENABLED=0 を指定する必要がある
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-w -s"

####################### Run stage #######################
FROM alpine:3.13
COPY --from=builder /app/cmd/go-rainy-bot/go-rainy-bot .
CMD ["./go-rainy-bot"]