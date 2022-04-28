# syntax=docker/dockerfile:1
ARG migrate

FROM golang:1.18-alpine
RUN apk add git
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
COPY *.go ./

RUN go build -o test-api .

EXPOSE 1324
#CMD ["go","run","."]
CMD ["./test-api"]