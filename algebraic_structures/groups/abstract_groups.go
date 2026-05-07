package groups

type Group[T any] interface {
	Op(a, b T) T
	Identity() T
	Inverse(a T) T
	Equals(a, b T) bool
	Order() int // if -1, too big to compute
	Family() string // not sure if I will keep this
	Contains(a T) bool
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

func (G GenericGroup[T]) Order() int {
	return G.order
}

func (G GenericGroup[T]) Family() string {
	return G.family
}
