package engine

import (
	"crypto/rand"
	"math/big"
)

// Returns an integer from 0 to the num - 1
func GetRandomInt(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64())
}

// Returns an integer from 1 to the num
func GetDiceRoll(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64()) + 1
}

// Returns a number between low and high, inclusive.
func GetRandomBetween(low, high int) int {
	return GetDiceRoll(high-low) + high
}
