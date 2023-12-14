# Stage 1: Build stage
FROM golang:latest AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

# Stage 2: Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/main .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]
