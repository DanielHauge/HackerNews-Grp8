package main

import (
	"log"
	"net/http"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"net"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
		//go RabbitMessage(Log_Q.Name, r.Method+","+r.RequestURI+","+name+","+time.Since(start).String())
		logz.Info(r.Method+","+r.RequestURI+","+name+","+time.Since(start).String())


	})
}

func SetupLogrus(){
	logger := logrus.New()
	conn, err := net.Dial("tcp", "ec2-18-216-94-144.us-east-2.compute.amazonaws.com:9600")
	if err != nil{
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "Go-Web-Api"}))
	if err!=nil{
		log.Fatal(err)
	}
	logger.Hooks.Add(hook)
	ctx := logger.WithFields(logrus.Fields{
		"method": "main",
	})
	ctx.Info("Logger Initialization Complete")
	logz = ctx
}

