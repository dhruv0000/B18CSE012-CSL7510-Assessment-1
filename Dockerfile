# Using alpine version since it has less image size
FROM alpine:latest

# Installing golang and its dependencies
RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin


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