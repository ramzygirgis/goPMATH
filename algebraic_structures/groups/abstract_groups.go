package groups

type Group[T any] interface {
	Op(a, b T) T
	Identity() T
	Inverse(a T) T
	Equals(a, b T) bool
}

type GenericGroup[T any] struct {
	op       func(T, T) T
	identity T
	inverse  func(T) T
	equals func(T, T) bool
	order int // 0 if infinite
	family string
}

func (G GenericGroup[T]) Op(a, b T) T {
	return G.op(a, b)
}

func (G GenericGroup[T]) Identity() T {
	return G.identity
}

func (G GenericGroup[T]) Inverse(a T) T {
	return G.inverse(a)
}

func (G GenericGroup[T]) Equals(a, b T) bool {
	return G.equals(a, b)
}
