# Build stage
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN cd cmd/server && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /server .
RUN pwd
WORKDIR /app
RUN ls -la
RUN pwd

# Run stage
FROM alpine:3.15
COPY --from=0 server .
RUN ls -la
RUN pwd
COPY .env .

EXPOSE 8080
CMD [ "/app/main" ]
