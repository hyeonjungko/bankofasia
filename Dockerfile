# Build binary file
FROM golang:1.20.5-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run binary file
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 8080

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
