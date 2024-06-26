# Dockerfile for Service 2
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY service2.go ./

RUN go build -o /service2-app

EXPOSE 8082

CMD [ "/service2-app" ]
