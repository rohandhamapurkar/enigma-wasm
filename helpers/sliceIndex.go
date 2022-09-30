package helpers

func StringIndexOf(element string, data string) int {
	for k, v := range data {
		if element == string(v) {
			return k
		}
	}
	return -1 //not found.
}

func SliceIndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
