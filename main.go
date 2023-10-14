package main

import (
	"fmt"
	"time"
	"math"
)

type  tokenBucket struct{
	tokens float64
	id string
	refillTime time.Time
	maxTokens float64
	lastRefillTime time.Time
	refillTimeRate float64
}

func NewTokenBucket(tokens float64,rate float64) *tokenBucket{
	return &tokenBucket{
		tokens: tokens,
		maxTokens: tokens,
		lastRefillTime: time.Now(),
		// id: id,
		refillTimeRate: rate,
	}
}

// function to refill the tokenBucket
func (tb* tokenBucket) refill(){
	currentTime := time.Now()
	timeSinceLastRefill := currentTime.Sub(tb.lastRefillTime)
	tokensToAdd :=  timeSinceLastRefill.Seconds() * tb.refillTimeRate
	tb.tokens = math.Min(tb.tokens+tokensToAdd, tb.maxTokens)
	tb.lastRefillTime = time.Now()
}

func (tb* tokenBucket) ratelimiter(requestNumber float64) bool{
	// refill the bucket based on refillTime 
	tb.refill()
	fmt.Println("add =",tb.tokens)
	if requestNumber < tb.tokens{
		tb.tokens -= requestNumber
		fmt.Println("sub = ",tb.tokens)
		return true
	}
	return false
}

func main(){
	newTokenBucket := NewTokenBucket(10,1)
	var i float64
	for  i=0; i<30; i++ {
		fmt.Println("rate limit check ", i+1, newTokenBucket.ratelimiter(20))
		time.Sleep(500 * time.Millisecond)
	} 
}

