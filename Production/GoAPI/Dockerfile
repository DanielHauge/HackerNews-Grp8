FROM golang:jessie

# Install any needed dependencies...
# RUN go get ...
RUN go get "github.com/go-sql-driver/mysql"
RUN go get github.com/gorilla/mux
RUN go get github.com/streadway/amqp
RUN go get github.com/rs/cors
RUN go get github.com/prometheus/client_golang/prometheus



# Copy the server code into the container
COPY Handlers.go Handlers.go
COPY Logger.go Logger.go
COPY Main.go Main.go
COPY MySQLHandler.go MySQLHandler.go
COPY RabbitManager.go RabbitManager.go
COPY PostClass.go PostClass.go
COPY Router.go Router.go
COPY Routes.go Routes.go

# Make port 8787 available to the host
EXPOSE 8787

RUN go build


