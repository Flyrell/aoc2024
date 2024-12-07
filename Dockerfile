FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY go.* ./
COPY .air.toml ./
RUN go install github.com/air-verse/air@latest
RUN go mod download -x && go mod verify

CMD ["air", "-c", ".air.toml"]
