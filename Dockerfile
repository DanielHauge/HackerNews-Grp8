FROM golang:jessie

# Install any needed dependencies...
# RUN go get ...
RUN go get gopkg.in/mgo.v2 

# Set the working directory
WORKDIR /src

# Copy the server code into the container
COPY basic_http_server.go basic_http_server.go

# Make port 8080 available to the host
EXPOSE 8866

# Build and run the server when the container is started
RUN go build /src/basic_http_server.go
ENTRYPOINT ./basic_http_server