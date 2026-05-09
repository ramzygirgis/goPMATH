package rings

type Polynomial[T any] struct {
	coefficients []T
	degree int
}



func (R Ring) PolynomialRing() Ring[list[T]] {

}

// either implement a polynomial struct
