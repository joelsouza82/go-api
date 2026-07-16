#FROM golang:1.26

# set working directory
#WORKDIR /go/src/app

# Copy the source code
#COPY . . 

#EXPOSE the port
#EXPOSE 8000

# Build the Go app
#RUN go build -o main cmd/main.go

# Run the executable
#CMD ["./main"]

ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
