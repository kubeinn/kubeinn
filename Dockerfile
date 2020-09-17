FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o . ./cmd/main.go
CMD ["/app/main"]