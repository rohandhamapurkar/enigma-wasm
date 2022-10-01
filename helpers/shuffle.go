package helpers

import (
	"math/rand"
	"time"
)

func ShuffleStringSlice(a []string) []string {
	result := make([]string, len(a))
	copy(result, a)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	
	return result
}
