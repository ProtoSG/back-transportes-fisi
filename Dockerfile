FROM golang:1.20 AS build

WORKDIR /app

COPY go.* ./ 

RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/main.go

EXPOSE 8080

COPY ./*.db ./

CMD ["/app/main"]
