package ratelimiter

import (
	"math"
	"time"
)

type TokenBucket struct {
	tokens         int
	refillTime     time.Time
	maxTokens      int
	lastRefillTime time.Time
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
	timeSinceLastRefill := currentTime.Sub(tb.lastRefillTime)
	tokensToAdd := timeSinceLastRefill.Seconds() * refillRate
	tb.tokens = int(math.Min(float64(tb.tokens)+tokensToAdd, float64(tb.maxTokens)))
	tb.lastRefillTime = time.Now()
}

func (tb *TokenBucket) IsRequestAllowed() bool {
	if tb.tokens > 0 {
		tb.tokens--
		return true
	} else {
		return false
	}
}
