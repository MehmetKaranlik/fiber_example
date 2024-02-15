package helper

func Map[T any, U any](arr []T, f func(T) U) []U {
	var result []U
	for _, v := range arr {
		result = append(result, f(v))
	}
	return result
}
