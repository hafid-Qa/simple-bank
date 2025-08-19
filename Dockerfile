# Build stage
FROM golang:1.23-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

# Install migrate CLI
RUN apk add --no-cache curl
# Corrected: Download the LINUX version
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz
# The extracted binary is now at /app/migrate.linux-amd64

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
# Corrected: Copy the extracted binary and rename it to 'migrate'
COPY --from=builder /app/migrate ./migrate
COPY .env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migrations ./db/migrations

# EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]