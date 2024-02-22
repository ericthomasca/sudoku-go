FROM golang:1.21

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /sudoku-go-api

EXPOSE 8976

CMD ["/sudoku-go-api"]