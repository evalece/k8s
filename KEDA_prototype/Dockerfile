# Stage 1: Compiler (large)
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o dataPull .

# Stage 2:  // use builder, discards Stage 1, hence slim image
FROM scratch
COPY --from=builder /app /app/dataPull
ENTRYPOINT ["/app/dataPull"]

# docker tag <IMAGE ID > datapull-app:latest