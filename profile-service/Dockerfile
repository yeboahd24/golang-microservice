# Use the official Golang image to create a build artifact.
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum and download dependencies.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the local packages and other necessary files into the container.
COPY . .

# Build the Go app.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o profile-service .

# Start a new stage from a smaller image.
FROM alpine:latest  

# Add certificates to enable HTTPS.
RUN apk --no-cache add ca-certificates

# Set the Working Directory inside the container.
WORKDIR /app

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/profile-service .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["./profile-service"]
