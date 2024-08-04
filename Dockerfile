FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . ./

RUN go build -o bin/goshort cmd/server/main.go

EXPOSE 8080

CMD [ "bin/goshort" ]
