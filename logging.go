package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

// #region Pre-recorded messages
const (
	callFailure string = "Echec de l'appel"
	callSuccess string = "Réussite de l'appel"
)

// #endregion Pre-recorded messages

func initMiddleware(router *mux.Router) {

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s\t%s\t", r.Method, r.RequestURI)
			printRequest(r.RemoteAddr)
			//constructHeaders(&w, r)
			next.ServeHTTP(w, r)

		})
	})

}

// LoggerHandler Middleware qui permet de logger les requêtes entrantes
func LoggerHandler(next http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		x, err := httputil.DumpRequest(r, true)
		if err != nil {
			failOnError(err, "An error occured during the request")
		}
		log.Printf("[BEGIN CALL] - %s\t%s\t", r.Method, r.RequestURI)
		log.Println("[HEADER] - ", fmt.Sprintf("%q", x))

		next.ServeHTTP(w, r)

		log.Printf(
			"[END CALL] - %s\t%s\t%s\t",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)

		w.WriteHeader(http.StatusOK)

	})
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("[ERROR] - %s: %s", msg, err)
	}
}

func printRequest(addr string) {
	log.Printf("[ %s ] - Request from %s ", time.Now().Format(time.RFC3339), addr)
}

func trace(desc string) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	log.Printf("%s - %s\n", frame.Function, desc)
}
