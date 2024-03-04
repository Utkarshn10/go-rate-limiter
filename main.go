package main

import (
	"fmt"
	"net/http"
)

func apiRequestHandler() {
	http.HandleFunc("/limited", limitedRequestHandler)
	http.HandleFunc("/unlimited", unlimitedRequestHandler)
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
