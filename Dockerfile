FROM golang:1.21

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go get github.com/ericthomasca/sudoku-go/sudoku

RUN go build -o /sudoku-go

EXPOSE 8976

CMD ["/sudoku-go"]