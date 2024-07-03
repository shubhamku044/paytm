package utils

import (
	"math/rand"
)

func GenerateRandomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}
