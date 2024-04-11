# Use the official Golang image from Docker Hub
FROM golang:latest

# Set a label for the description
LABEL description="Web Simple Application"

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o web .

# Expose port 8080 to the outside world
EXPOSE 8080

# Define the entry point for the container
ENTRYPOINT [ "./web" ]
