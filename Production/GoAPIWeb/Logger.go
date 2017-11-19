package main

import (
	"log"
	"net/http"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"net"
	"io"

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
// New logging https://www.goinggo.net/2013/11/using-log-package-in-go.html
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)
// Examples
// Trace.Println("I have something standard to say")
// Info.Println("Special Information")
// Warning.Println("There is something you need to know about")
// Error.Println("Something has failed")


// 	Example LogSetup(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

func LogSetup(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
// END New logging

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




