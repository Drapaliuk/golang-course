package utils

func Delete[T any](s []T, idx int) []T {
	if idx < 0 || idx >= len(s) {
		return s
	}

	return append(s[:idx], s[idx+1:]...)
}

func Filter[T any](s []T, cb func(el T, i int) bool) []T {
	res := []T{}

	for i, el := range s {
		if cb(el, i) {
			res = append(res, el)
		}
	}

	return res
}

func FindIndex[T any](s []T, cb func(el T, i int) bool) int {
	for i, el := range s {
		if cb(el, i) {
			return i
		}
	}

	return -1
}

func Find[T any](s []T, cb func(el T, i int) bool) *T {
	for i, el := range s {
		if cb(el, i) {
			return &el
		}
	}

	return nil
}

func Some[T any](s []T, cb func(el T, i int) bool) bool {
	for i, el := range s {
		if cb(el, i) {
			return true
		}
	}

	return false
}
