FROM golang:1.24-alpine as builder

# Creates an application directory to hold your application’s source code
WORKDIR /app

RUN apk add --no-cache git

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .

RUN go build -o /training-plan cmd/api/main.go

FROM alpine:latest as runner

WORKDIR /

COPY --from=builder /training-plan /training-plan

# Specifies the executable command that runs when the container starts
CMD ["/training-plan"]
