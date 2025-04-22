# Gunakan image Golang sebagai base image
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

ENV PATH="/go/bin:${PATH}"

COPY docker/entrypoint.sh /docker/entrypoint.sh
RUN chmod +x /docker/entrypoint.sh

COPY go.mod go.sum ./
RUN go mod download

# Copy semua file ke container
COPY . .

RUN go build -o /app/server ./main.go

FROM golang:1.23-alpine AS final
WORKDIR /app

COPY --from=builder /app/server /app/server
COPY --from=builder /docker/entrypoint.sh /docker/entrypoint.sh
COPY .env.development .env

EXPOSE 5000

ENTRYPOINT ["/docker/entrypoint.sh"]

