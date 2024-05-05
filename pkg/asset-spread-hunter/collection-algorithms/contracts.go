package collection_algorithms

type Queue[T any] interface {
	Push(value T)
	Pop() (*T, error)
	Size() int64
	IsEmpty() bool
}
