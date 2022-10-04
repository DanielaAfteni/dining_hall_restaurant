package main

import (
	// Package json implements encoding and decoding of JSON.
	// The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.

	// Package ioutil implements some I/O utility functions.

	// Package os provides a platform-independent interface to operating system functionality.

	// importing the gin, because is a high-performance HTTP web framework written in Golang (Go).

	"github.com/rs/zerolog/log"
)

var Ratings = [6]float32{0, 0, 0, 0, 0, 0}

func GetRating(maxWait float32, cookingTime int64) {
	var currentRating int
	fCookingTime := float32(cookingTime)
	fMaxWait := float32(maxWait)
	if fCookingTime < fMaxWait {
		Ratings[5]++
		currentRating = 5
	} else if fCookingTime < fMaxWait*1.1 {
		Ratings[4]++
		currentRating = 4
	} else if fCookingTime < fMaxWait*1.2 {
		Ratings[3]++
		currentRating = 3
	} else if fCookingTime < fMaxWait*1.3 {
		Ratings[2]++
		currentRating = 2
	} else if fCookingTime < fMaxWait*1.4 {
		Ratings[1]++
		currentRating = 1
	} else {
		Ratings[0]++
		currentRating = 0
	}
	log.Printf("New rating from served clients = %d", currentRating)
	CalculateFinalRating()
}
