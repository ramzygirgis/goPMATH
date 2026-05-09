package rings

type Ring[T any] interface {
	Add(a, b T) T
	Zero() T
	AddInverse(a T) T
	Mult(a, b) T
	One(a, b) T
	Equals(a, b T) bool
	Contains(a T) bool
}
