FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o statefully .

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/statefully /app/statefully
CMD ["./statefully"]
