FROM golang:jessie

# Install any needed dependencies...
# RUN go get ...

# Copy the server code into the container
COPY server.go server.go

# Make port 8080 available to the host
EXPOSE 8081

# Build and run the server when the container is started
ENTRYPOINT go run server.go
