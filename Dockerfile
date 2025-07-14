# syntax=docker/dockerfile:1

# =========================
# STEP 1 - BUILD
# =========================
FROM golang:1.23.8-alpine AS builder

WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main ./cmd/server

# =========================
# STEP 2 - DEPLOY
# =========================
FROM alpine:latest

# Set environment variables for user/group
ENV UID=1001
ENV USER=appuser
ENV GID=1001
ENV GROUP=appgroup

# Create non-root user and group
RUN addgroup -g $GID -S $GROUP && \
    adduser -S -u $UID -G $GROUP $USER

# Install ca-certificates and timezone
RUN apk add --no-cache ca-certificates tzdata && update-ca-certificates

# Set the timezone
ENV TZ=Asia/Ho_Chi_Minh
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy config and protogen if needed at runtime
COPY --from=builder /app/config ./config
COPY --from=builder /app/protogen ./protogen

# Change ownership to non-root user
RUN chown -R $USER:$GROUP /app

# Expose the service port
EXPOSE 50054

# Run the application as non-root user
USER $USER

# Start the application
ENTRYPOINT ["./main"]