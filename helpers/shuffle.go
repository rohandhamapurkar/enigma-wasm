package helpers

import (
	"math"
	"math/rand"
)

func ShuffleStringSlice(a []string) []string {
	result := make([]string, len(a))
	copy(result, a)
	for i := len(result) - 1; i > 0; i++ {
		randomIndex := int(math.Floor(rand.Float64() * float64(i + 1)))
		temp := result[i]
		result[i] = result[randomIndex]
		result[randomIndex] = temp
	}
	return result
}
