# Start from official Go image
FROM golang:1.24

# Set working directory
WORKDIR /app

# Copy files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go binary
RUN go build -o main .

# Run the binary
CMD ["./main"]
