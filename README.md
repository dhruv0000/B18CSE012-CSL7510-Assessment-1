# B18CSE012-CSL7510-Assessment-1

This repo contains an basic web application build from scratch to insert data record in a db, to be used in `Assessment 1: Virtual Machine and Dockers` of CSL7510.

## Part 1

https://youtu.be/hw9YdIAC7Ow

## Part 2

### Dockerfile

The dockerfile has 12 steps discussed below:

1. ```dockerfile
    FROM golang:1.16
    ```

    We are creating our image using the base image golang:1.16. 

2. ```dockerfile
   WORKDIR /src
   ```

   Move to working directory /src where we will copy our codebase

3. ```dockerfile
    COPY go.mod ./
    COPY go.sum ./
    RUN go mod download
   ```

   We copy the dependency files to the container and download the golang dependencies.

4. ```dockerfile
    COPY . ./
   ```

   Copy the rest of codebase to the working directory.

5. ```dockerfile
    RUN go build -o /main main.go
   ```

   Build the application and save the binary at `/main`

7. ```dockerfile
    EXPOSE 8080
   ```

   The app uses port 8080, so we expose that port from inside the container to the outside.

7. ```dockerfile
    CMD ["/main"]
   ```

   Execute the binary when the we run the image.

### Building the Image

```bash
docker build -t <iamge-name> .
```

### Running the image

```bash
docker run -p 8080:8080 ass1.1
```

NOTE: -p flag is used to publish port inside the container to port outside the container (`-p [host_port]:[container_port]`)

You can access the applaction at `localhost:8080/assets/index.html`.
