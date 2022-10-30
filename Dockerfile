# Build stage
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN cd cmd/server && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /server .
WORKDIR /app

# Run stage
FROM alpine:3.15
COPY --from=0 server .
COPY .env .

EXPOSE 8080
CMD [ "/app/main" ]
