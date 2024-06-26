# Dockerfile for Service 1
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY service1.go ./

RUN go build -o /service1-app

EXPOSE 8081

CMD [ "/service1-app" ]
