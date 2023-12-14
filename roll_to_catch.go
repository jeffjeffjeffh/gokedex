package main

import (
	"math/rand"
)

func rollToCatch(difficulty int) bool {
	numRolls := difficulty / 100

	for i := 0; i < numRolls; i++ {
		min := 1
		max := 101
		roll := rand.Intn(max - min) + min
		if roll < 50 {
			return false
		}
	}
	
	return true
}