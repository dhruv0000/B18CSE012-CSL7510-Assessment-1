# Using alpine version since it has less image size
FROM golang:1.16

# Move to working directory /src
WORKDIR /src

# Copy and download dependency using go mod
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the code into the working directory
COPY . ./

# Build the application
RUN go build -o /main main.go

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/main"]