package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var serveMux *http.ServeMux

func main() {
	initConfiguration()
	prepareLogs()
	serveMux = http.NewServeMux()
	router = NewRouter()
	initMiddleware(router)
	go serverHttpServer()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, os.Kill)

	select {
	case <-sigChan:
		time.Sleep(5 * time.Second)
		logFile.Close()
		os.Exit(0)
	}
}

func serverHttpServer() {
	log.Println("Serving at ", appConfig.httpListenUri(), "")
	log.Println(http.ListenAndServe(appConfig.httpListenUri(), router))
}
