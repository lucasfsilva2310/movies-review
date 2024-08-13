FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY .air.toml .

RUN go install github.com/air-verse/air@latest

EXPOSE 8000

ENTRYPOINT ["air", "-c", ".air.toml"]