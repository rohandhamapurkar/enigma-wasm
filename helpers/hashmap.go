package helpers

import (
	"golang-wasm/constants"
	"math"
	"math/rand"
)

func GetRandomCharacterHashMap(a []string) map[string]string {

	m := make(map[string]string, len(constants.Characters))
	set := make(map[string]bool, len(constants.Characters))
	for _, i := range a {
		found := false

		for !found {
			pair := a[int(math.Floor(rand.Float64()*float64(len(a))))]

			v1, ok1 := m[i]
			v2, ok2 := m[pair]
			ok3 := set[pair]
			if ok1 {
				m[v1] = i
				set[i] = true
				found = true
			} else if i == pair || ok3 {
				found = false
			} else if ok2 {
				m[v2] = pair
				set[pair] = true
				found = true
			} else {
				m[i] = pair
				m[pair] = i
				set[i] = true
				set[pair] = true
				found = true
			}
		}
	}
	return m
}
