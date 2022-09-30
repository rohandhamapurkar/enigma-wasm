package helpers

func SliceRotateRight(a []string) []string {
	result := make([]string, len(a))
	rlen := len(result)
	copy(result, a)
	last := result[rlen-1]

	return append([]string{last}, result[:rlen-1]...)
}