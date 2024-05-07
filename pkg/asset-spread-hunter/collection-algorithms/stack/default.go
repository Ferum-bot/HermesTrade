package stack

import (
	"errors"
	collectionalgorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/collection-algorithms"
)

var ErrorEmptyStack = errors.New("stack is empty")

type defaultStack[T any] struct {
	array []T
}

func NewDefaultStack[T any]() collectionalgorithms.CopyableStack[T] {
	return &defaultStack[T]{
		array: make([]T, 0),
	}
}

func (d *defaultStack[T]) Push(value T) {
	d.array = append(d.array, value)
}

func (d *defaultStack[T]) Pop() (*T, error) {
	if len(d.array) == 0 {
		return nil, ErrorEmptyStack
	}

	size := len(d.array)
	value := d.array[size-1]
	d.array = d.array[:size-1]

	return &value, nil
}

func (d *defaultStack[T]) Size() int64 {
	return int64(len(d.array))
}

func (d *defaultStack[T]) IsEmpty() bool {
	return len(d.array) == 0
}

func (d *defaultStack[T]) MakeCopy() collectionalgorithms.Copyable {
	arrayCopy := make([]T, len(d.array))
	for i, value := range d.array {
		arrayCopy[i] = value
	}

	return &defaultStack[T]{
		array: arrayCopy,
	}
}
