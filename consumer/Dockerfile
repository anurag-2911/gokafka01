FROM golang:1.22-bullseye

WORKDIR /app

# Copy module definitions first 
COPY go.mod .
COPY go.sum .

# Install GCC and other build tools using apt-get
RUN apt-get update && apt-get install -y build-essential

RUN go mod download

# Enable CGO when building
ENV CGO_ENABLED=1

COPY . .

RUN go build -o consumer
CMD ["./consumer"]
