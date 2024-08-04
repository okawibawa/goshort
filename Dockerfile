FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/goshort cmd/server/main.go

EXPOSE 8081

CMD [ "./bin/goshort" ]
