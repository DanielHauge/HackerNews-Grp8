FROM golang:jessie

# Install any needed dependencies...
# RUN go get ...
RUN go get "github.com/go-sql-driver/mysql"
RUN go get github.com/gorilla/mux

# Set the working directory
WORKDIR /src

# Copy the server code into the container
COPY Handlers.go Handlers.go
COPY Logger.go Logger.go
COPY main.go main.go
COPY MySQLHandler.go MySQLHandler.go
COPY PostClass.go PostClass.go
COPY Router.go Router.go
COPY Routes.go Routes.go
COPY Todo.go Todo.go

# Make port 8787 available to the host
EXPOSE 8787

# Build and run the server when the container is started
RUN go build /src/.
ENTRYPOINT ./GoAPI
