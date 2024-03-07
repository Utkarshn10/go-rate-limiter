# go-rate-limiter

Rate Limiter implemented using Token Bucket Algorithm written in Go.

# Overview
The token bucket algorithm allows a certain number of requests to be made in a given time period. It works by refilling a bucket with a fixed number of tokens at a constant rate. Each request consumes one token from the bucket. If the bucket is empty, no more requests can be made until it is refilled.

# Implementation
The TokenBucket struct represents the rate limiter. It contains the following fields:

- `tokens`: The current number of tokens in the bucket.
- `refillTime`: The time at which the bucket was last refilled.
- `maxTokens`: The maximum number of tokens the bucket can hold.
- `lastRefillTime`: The time at which the bucket was last refilled.
- `mutex`: A mutex to ensure thread safety.

# The package provides the following functions:

- `NewTokenBucket(tokens int) *TokenBucket`: Initializes a new TokenBucket with the specified number of tokens.
- `Refill()`: Refills the TokenBucket with tokens based on the refill rate.
- `IsRequestAllowed() bool`: Checks if a request is allowed based on the current number of tokens in the bucket.

# Sample usecase

```Golang
package main

import (
	"fmt"
	"time"

	"github.com/Utkarshn10/go-rate-limiter/src/ratelimiter"
)

func main() {
	// Create a new TokenBucket with 10 tokens
	tb := ratelimiter.NewTokenBucket(10)

	// Periodically refill the bucket
	go func() {
		for {
			tb.Refill()
			time.Sleep(1 * time.Second)
		}
	}()

	// Make requests and check if they are allowed
	for i := 0; i < 15; i++ {
		if tb.IsRequestAllowed() {
			fmt.Println("Request", i+1, "allowed")
		} else {
			fmt.Println("Request", i+1, "blocked")
		}
		time.Sleep(100 * time.Millisecond)
	}
}

```

