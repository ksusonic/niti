package utils

// Ptr returns a pointer to the given value.
func Ptr[T any](v T) *T {
	return &v
}

func NilIfEmpty[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}
