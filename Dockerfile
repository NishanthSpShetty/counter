FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod download

RUN go build -o ./counter

EXPOSE 8081

CMD [ "./counter" ]

