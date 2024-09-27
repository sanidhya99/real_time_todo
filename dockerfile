# Step 1: Use official Golang image as a base (updated to 1.20)
FROM golang:1.20-alpine AS build

# Step 2: Set environment variables
ENV GO111MODULE=on

# Step 3: Set the working directory inside the container
WORKDIR /app

# Step 4: Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Step 5: Download the dependencies
RUN go mod download

# Step 6: Copy the rest of the application code to the container
COPY . .

# Step 7: Build the Go app
RUN go build -o main .

# Step 8: Use a smaller image to run the app
FROM alpine:latest

WORKDIR /root/

# Copy binary from the build image
COPY --from=build /app/main .

# Expose the port that the app listens on (default for Gin: 8080)
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]