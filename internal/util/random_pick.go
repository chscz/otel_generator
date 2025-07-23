package util

import "math/rand"

func PickRandomElementFromSlice[T string | int](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	return slice[rand.Intn(len(slice))], true
}
