package interfaces

type RequestBody[T interface{}] interface {
	ValidatedParse(data []byte) (*T, error)
}
