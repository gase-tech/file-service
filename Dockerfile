FROM golang:1.15.3-alpine AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -v src/main.go
CMD ["/app/main"]