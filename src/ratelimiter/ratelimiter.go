package ratelimiter

import (
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens         int
	refillTime     time.Time
	maxTokens      int
	lastRefillTime time.Time
	mutex          sync.Mutex
}

var refillRate float64 = 1

func NewTokenBucket(tokens int) *TokenBucket {
	return &TokenBucket{
		tokens:         tokens,
		maxTokens:      tokens,
		lastRefillTime: time.Now(),
	}
}

// function to refill the TokenBucket
func (tb *TokenBucket) Refill() {
	currentTime := time.Now()
	tb.mutex.Lock()
	timeSinceLastRefill := currentTime.Sub(tb.lastRefillTime)
	tokensToAdd := timeSinceLastRefill.Seconds() * refillRate
	tb.tokens = int(math.Min(float64(tb.tokens)+tokensToAdd, float64(tb.maxTokens)))
	tb.lastRefillTime = time.Now()

	tb.mutex.Unlock()
}

// checks if the request can be made or not. If tokens are greater than 0 then return true
func (tb *TokenBucket) IsRequestAllowed() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	} else {
		return false
	}
}
