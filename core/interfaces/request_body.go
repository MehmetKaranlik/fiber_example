package interfaces

type RequestBody[T RequestBody[T]] interface {
	ValidatedParse(data []byte) (*T, error)
}
