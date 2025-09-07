# Stage 1: Build
FROM golang:1.24.6-alpine AS builder

WORKDIR /server

# Cache dependencies
COPY go.mod ./
RUN go mod download

# Copy source
COPY data/ ./data
COPY src/ .

# Build static binary
RUN CGO_ENABLED=0 go build -o server

# Stage 2: Minimal runtime
FROM scratch

WORKDIR /server

# Copy only the final binary
COPY --from=builder /server/server .

EXPOSE 8080

CMD ["./server"]
