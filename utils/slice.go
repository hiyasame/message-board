package utils

func Map[T interface{}, V interface{}](slice []T, mapper func(T) V) []V {
	res := make([]V, len(slice))
	for k, v := range slice {
		res[k] = mapper(v)
	}
	return res
}
