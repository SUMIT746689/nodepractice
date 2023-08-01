FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# COPY . .

RUN go mod tidy

# RUN CGO_ENABLED=0 go build -o api_server cmd/api_server/main.go

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

ENTRYPOINT ["air", "-c", "api_server.air.toml"]
