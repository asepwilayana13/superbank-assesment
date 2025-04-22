FROM golang:1.24-alpine AS builder

WORKDIR /app
ENV PATH="/go/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY docker/entrypoint.sh /docker/entrypoint.sh
RUN chmod +x /docker/entrypoint.sh

RUN go build -o /app/server ./main.go

FROM alpine:latest AS final

WORKDIR /app

COPY --from=builder /app/server ./
COPY --from=builder /docker/entrypoint.sh ./docker/entrypoint.sh
COPY .env.development .env

RUN chmod +x ./docker/entrypoint.sh
RUN ls -l ./docker/entrypoint.sh && cat ./docker/entrypoint.sh


EXPOSE 8000

ENTRYPOINT ["./docker/entrypoint.sh"]
