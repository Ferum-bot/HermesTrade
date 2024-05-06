package stack

import collection_algorithms "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/collection-algorithms"

type defaultStack[T any] struct {
	array []T
}

func NewDefaultStack[T any]() collection_algorithms.CopyableStack[T] {
	return &defaultStack[T]{}
}

func (d *defaultStack[T]) Push(value T) {
	//TODO implement me
	panic("implement me")
}

func (d *defaultStack[T]) Pop() (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (d *defaultStack[T]) Size() int64 {
	//TODO implement me
	panic("implement me")
}

func (d *defaultStack[T]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (d *defaultStack[T]) MakeCopy() collection_algorithms.Copyable {
	//TODO implement me
	panic("implement me")
}
