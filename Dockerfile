FROM golang:1.21 AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o sudoku-go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/sudoku-go .

EXPOSE 8976

CMD ["./sudoku-go"]
