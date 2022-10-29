# Build stage
FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go get google.golang.org/grpc/internal/channelz@v1.47.0
RUN go build -o main main.go

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
COPY start.sh .
COPY wait-for.sh .
RUN addgroup --system user && adduser --system -G user user && \
    chown user:user -R . && chmod -R 755 . && \
    chmod +x start.sh && chmod +x wait-for.sh
RUN chmod +x start.sh && chmod +x wait-for.sh
COPY db/migration ./db/migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
