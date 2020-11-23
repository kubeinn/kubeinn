FROM golang:alpine3.12
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o . ./cmd/main.go
CMD ["/app/main"]