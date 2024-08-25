package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Random3DigitString() string {
	rand.Seed(time.Now().UnixNano())     // Seed the random number generator
	randomNumber := rand.Intn(900) + 100 // Generate a random number between 100 and 999
	return fmt.Sprintf("%03d", randomNumber)
}
