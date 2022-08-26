package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "OK")
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)

	router.GET("/health", HealthCheck)

	log.Fatal(http.ListenAndServe(":8081", router))
}
