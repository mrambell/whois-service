# Use the official Golang image as a base
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o whois-service

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./whois-service"]