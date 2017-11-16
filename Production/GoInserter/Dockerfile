FROM golang:jessie

# Install any needed dependencies...
# RUN go get ...
RUN go get "github.com/go-sql-driver/mysql"
RUN go get github.com/streadway/amqp
RUN go get github.com/gin-gonic/gin/json

# Copy the server code into the container
COPY Main.go Main.go
COPY MySQLHandler.go MySQLHandler.go
COPY RabbitManager.go RabbitManager.go
COPY Data.go Data.go

RUN go build

