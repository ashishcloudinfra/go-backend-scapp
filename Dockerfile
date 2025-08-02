# Build stage: compile the Go application using Alpine-based Go image.
FROM golang:1.24-alpine AS builder

# Set the working directory for building the app.
WORKDIR /app

# Copy all source files into the build container.
COPY . .

# Download Go module dependencies.
RUN go mod download

# Build the Go application into an executable (named "myapp").
RUN go build -o myapp .

# Final stage: Create a lightweight runtime container based on Alpine.
FROM alpine:3.21

# Install Python 3, pip, bash, and any runtime dependencies.
RUN apk update && \
    apk add --no-cache python3 py3-pip bash

# Install the required Python package globally.
RUN pip3 install --no-cache-dir --break-system-packages yfinance

# Set the working directory for the runtime container.
WORKDIR /app

# Copy the built Go binary from the builder stage.
COPY --from=builder /app/myapp .

# Copy only the scripts folder from the builder stage.
COPY --from=builder /app/scripts /app/scripts

# Set an environment variable (adjust as needed, for example, to force production branch).
ENV ENV=production

# Expose the HTTP port your API is listening on.
EXPOSE 8080

# Run the Go application.
CMD ["./myapp"]
