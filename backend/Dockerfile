FROM golang:latest

# Set the working directory to the root of the app
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the source code to the working directory
COPY . .

# Build the app
RUN go build -o backend .

# Expose the port that the app will listen on
EXPOSE 8080

# Run the app
CMD ["./backend"]

