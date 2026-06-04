package rings

type Polynomial[T any] struct {
	ring Ring[T] // can prob remove
	coefficients []T
}


type PolynomialRing[T any] struct { // this will implement Ring[Polynomial[T]]
	ring Ring[T]
}


// constructor:
func (R Ring) PolynomialRing() PolynomialRing[T] {
	return PolynomialRing[T]{ring: R}
}



func (R PolynomialRing[T]) Zero[T any]() T {
	ring_zero := R.ring.Zero()
	return Polynomial[T]{ring: R, coefficients: []T{ring_zero}}  // TODO: Remove coeffs field
}


func (R PolynomialRing[T]) Add[T any] (f, g Polynomial[T]) Polynomial[T] {
	deg_f := len(f.coefficients) - 1
	deg_g := len(g,coefficients) - 1
	max_len := max(deg_f + 1, deg_g + 1)
	sum_coeffs := make([]T, max_len)
	for i := 0; i < max_len; i++ {
		sum_coeffs[i] = f.coefficients[i] + g.coefficients[i] 
	}
	leading_index = max_len - 1
	while sum_coeffs[leading_index] == 0 {
		if leading_index == 0 {
			return R.Zero[T]()
		}
		leading_index -= 1
	}
	coeffs = sum_coeffs[:leading_index + 1]
	return Polynomial[T]{ring: R, coefficients: coeffs} // TODO: Remove coeffs field
	}
}


func (R PolynomialRing[T]) AddInverse[T any](f Polynomial[T]) {
	new_coeffs := make([]T, len(f.coefficients))
	var coeff T
	for i := 0; i < len(f.coefficients); i++ {
		coeff = f.coefficients[i]
		inverse = R.ring.AddInverse(coeff)
		new_coeffs[i] = inverse
	}
	return Polynomial[T]{ring: R, coefficients: new_coeffs} // TODO: Remove coeffs field
}



func (R PolynomialRing[T]) Degree[T any](f Polynomial[T]) { // does not need to be a ring method, can be f.degree() or smth instead
	return len(f.coefficients) - 1
}



// either implement a polynomial struct
