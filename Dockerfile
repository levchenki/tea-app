FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tea-bot-api ./cmd/tea-app/main.go

RUN chmod +x tea-bot-api

CMD [ "./tea-bot-api" ]