package collection_algorithms

type Queue[T any] interface {
	Push(value T)
	Pop() (*T, error)
	Size() int64
	IsEmpty() bool
}

type Stack[T any] interface {
	Push(value T)
	Pop() (*T, error)
	Size() int64
	IsEmpty() bool
}

type Copyable interface {
	MakeCopy() Copyable
}

type CopyableStack[T any] interface {
	Stack[T]
	Copyable
}
