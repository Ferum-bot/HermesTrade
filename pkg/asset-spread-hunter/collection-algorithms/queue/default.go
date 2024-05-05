package queue

import (
	"errors"
	collection_algorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/collection-algorithms"
)

var ErrorEmptyQueue = errors.New("queue is empty")

type defaultQueue[T any] struct {
	array []T
}

func NewDefaultQueue[T any]() collection_algorithms.Queue[T] {
	return &defaultQueue[T]{
		array: make([]T, 0),
	}
}

func (d *defaultQueue[T]) Push(value T) {
	d.array = append(d.array, value)
}

func (d *defaultQueue[T]) Pop() (*T, error) {
	if len(d.array) < 1 {
		return nil, ErrorEmptyQueue
	}

	element := d.array[0]
	d.array = d.array[1:]
	return &element, nil
}

func (d *defaultQueue[T]) Size() int64 {
	return int64(len(d.array))
}

func (d *defaultQueue[T]) IsEmpty() bool {
	return len(d.array) == 0
}
