package groups

type Group[T any] interface {
	Op(a, b T) T
	Identity() T
	Inverse(a T) T
}

type GenericGroup[T any] struct {
	op       func(T, T) T
	identity T
	inverse  func(T) T
}

func (g GenericGroup[T]) Op(a, b T) T {
	return g.op(a, b)
}

func (g GenericGroup[T]) Identity() T {
	return g.identity
}

func (g GenericGroup[T]) Inverse(a T) T {
	return g.inverse(a)
}
