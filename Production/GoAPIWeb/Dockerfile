FROM golang:jessie

# Install any needed dependencies...
# RUN go get ...
RUN go get "github.com/go-sql-driver/mysql"
RUN go get github.com/gorilla/mux
RUN go get github.com/streadway/amqp
RUN go get github.com/rs/cors
RUN go get -u gopkg.in/russross/blackfriday.v2
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/sirupsen/logrus
RUN go get github.com/bshuster-repo/logrus-logstash-hook


# Copy the server code into the container
COPY Handlers.go Handlers.go
COPY Logger.go Logger.go
COPY Main.go Main.go
COPY MySQLHandler.go MySQLHandler.go
COPY RabbitManager.go RabbitManager.go
COPY Data.go Data.go
COPY Router.go Router.go
COPY Routes.go Routes.go
COPY GmailMailAPI.go GmailMailAPI.go

# Make port 9191 available to the host
EXPOSE 9191

RUN go build
