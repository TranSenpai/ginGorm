# Multi-stage builds
# In a multi-stage Docker build,
# each stage can use the build artifacts (the results like compiled binaries, libraries) from the previous stages,
# but the final image will only contain what you manually copy from previous stages —
# and nothing else like compilers or source code!
# The last image layer in multi-stage build is a complete program,
# -> but it only has the final built binary + minimal files needed to run,
# -> without all the extra build tools (Go compiler, caches, source code, etc.).

# In short:
# "In multi-stage builds, we separate building and running:
# the final Docker image contains only the built program, no build tools."

# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files first for caching
# Lists which packages (modules) your project depends on (and their versions).
# Verifies the integrity (hashes) of those packages to ensure they have not been tampered (handle) with.
COPY go.mod go.sum ./

# Download dependencies early
RUN go mod download

COPY . .

# Build the binary
# The tag-o always consistent name, no matter what the folder name is.
# Normally, when run go build. It will generate a binary with the same name as folder.
# Use this tag to force it's name
# CGO_ENABLED=0 Turn OFF CGO. Build pure Go code. No linking C libraries like libc, libmysqlclient, etc.
# Set output binary for Linux, no matter what platform container built on (Windows, Mac, Linux).
# It is because the Docker container run in linux environment.
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Run
FROM alpine:latest

# # Install certificates if using HTTPS in your app later
# # apl = Alpine Linux Package Manager
# When your Go app or any client connects to an HTTPS server (like https://api.example.com)
#  It needs to verify the server’s SSL certificate, 
# And to do that, it needs a set of trusted Certificate Authority (CA) certificates (like Let's Encrypt, DigiCert, etc.).
# RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Expose application port
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
