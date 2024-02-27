FROM golang:alpine

WORKDIR /app

# Copy module definitions first 
COPY go.mod .
COPY go.sum .

RUN go mod download  # Ensure modules are downloaded

# Enable CGO when building
ENV CGO_ENABLED=1

COPY . .

RUN go build -o consumer 
CMD ["./consumer"]
