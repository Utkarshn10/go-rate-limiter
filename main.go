package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Utkarshn10/go-rate-limiter/src/ratelimiter"
)

// this is an example of how go-rate-limiter works
var ip_addresses = make(map[string]*ratelimiter.TokenBucket)

func apiRequestHandler() {
	http.HandleFunc("/limited", func(w http.ResponseWriter, r *http.Request) {
		ip_addr := r.RemoteAddr

		_, ok := ip_addresses[ip_addr]
		if !ok {
			ip_addresses[ip_addr] = ratelimiter.NewTokenBucket(10)
		}

		if ip_addresses[ip_addr].IsRequestAllowed() {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Limited")
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprint(w, "Too many requests!\n")
		}
	})

	go func() {
		for {
			if len(ip_addresses) > 0 {
				time.Sleep(time.Second)
				for _, tb := range ip_addresses {
					tb.Refill()
				}
			}
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	apiRequestHandler()
}
