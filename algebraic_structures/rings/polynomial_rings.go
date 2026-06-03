package rings

type Polynomial[T any] struct {
	ring Ring[T]
	coefficients []T
	// degree int // maybe remove
}

func (R Ring[Polynomial[T]]) Zero[T any]() T {
	ring_zero := R.ring.Zero()
	return Polynomial[T]{ring: R, coefficients: []T{ring_zero}} 
}

func (R Ring[Polynomial[T]]) Add[T any] (f, g Polynomial[T]) Polynomial[T] {
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
	return Polynomial[T]{ring: R, coefficients: coeffs}
	}
}


func (R Ring[Polynomial[T]]) Degree[T any](f Polynomial[T]) {
	return len(f.coefficients) - 1
}


func (R Ring) PolynomialRing() Ring[[]list] {

}

// either implement a polynomial struct
