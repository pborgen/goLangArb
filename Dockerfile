# syntax=docker/dockerfile:1

FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

#Copy Everything
COPY . .

RUN ls

WORKDIR /app/cmd

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o start /app/cmd


# Run
CMD ["/app/cmd/start"]