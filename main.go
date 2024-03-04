package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Utkarshn10/go-rate-limiter/src/ratelimiter"
)

func apiRequestHandler() {
	tokenBucketRateLimiter := ratelimiter.NewTokenBucket(10, 1, nil)
	http.HandleFunc("/limited", func(w http.ResponseWriter, r *http.Request) {
		if tokenBucketRateLimiter.Ratelimiter(1) {
			fmt.Fprint(w, "Limited")
		} else {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		}
	})
	http.HandleFunc("/unlimited", func(w http.ResponseWriter, r *http.Request) {
		if tokenBucketRateLimiter.Ratelimiter(1) {
			fmt.Fprint(w, "Unlimited")
		} else {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	apiRequestHandler()
}
