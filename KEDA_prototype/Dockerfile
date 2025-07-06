
# Two stage build
# Stage 1; Build stage: Building with SDk and language specific tools to complete compilation of source code.
# takes source code as input and outputs binary code of given Arch dependency (i.e., apline)
FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY . .
RUN go build -ldflags="-s -w" -o dataPull .

# Stage 2; Run stage:  // Runs the binary code generated from stage 1. We will export this stage as slim Docker image 
FROM scratch
COPY --from=builder /app/dataPull /app/dataPull
ENTRYPOINT ["/app/dataPull"]

# docker tag <IMAGE ID > datapull-app:latest

# 
EXPOSE 8066 