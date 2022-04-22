# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
RUN apk add git
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
COPY *.go ./

RUN go build -o test-api2 .

EXPOSE 1324

CMD ["go","run","."]
#CMD ["./test-api"]