FROM golang:latest

RUN go get "github.com/joho/godotenv"
RUN go get "github.com/benmanns/goworker"
RUN go get "github.com/lib/pq"

RUN mkdir /go/src/app
WORKDIR /go/src/app

COPY . .

RUN go build -o main .

CMD ["./main"]
