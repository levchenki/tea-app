FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o tea-app ./cmd/tea-app/main.go

RUN chmod +x tea-app

CMD [ "./tea-app" ]