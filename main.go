package main

import (
	"fmt"
	"log"
	"net/http"
)

func apiRequestHandler() {
	http.HandleFunc("/limited", limitedRequestHandler)
	http.HandleFunc("/unlimited", unlimitedRequestHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func limitedRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Limited")
}

func unlimitedRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Unlimited")
}

func main() {
	apiRequestHandler()
}
