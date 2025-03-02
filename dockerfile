# Step 1: Build the Go binary in a temporary container
# Using the official Go image as the build environment
FROM golang:latest AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum haven't changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /my-go-app

# Step 2: Create the final container with only the necessary files
# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the previous stage
COPY --from=build /my-go-app .

# Expose port if your Go app uses a specific port (e.g., for a web server)
# EXPOSE 8080

# Command to run the Go app
CMD ["./my-go-app"]
# CMD ["sleep", "infinity"]