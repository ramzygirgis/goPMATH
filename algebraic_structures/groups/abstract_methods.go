package groups


func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func Pow[T any](G Group[T], g T, n int) T {
	res := G.Identity()
	
	base := g
	if n < 0 {
		base = G.Inverse(g)
	}
	
	for i := 0; i < abs(n); i++ {
		res = G.Op(res, base)
	}
}

func (G Group[T]) Commutator[T any](a T, b T) T {
	ab := G.Op(a, b)
	ba := G.Op(b, a)
	ba_inverse = G.Inverse(ba)
	return G.Op(ab, ba_inverse)
}

func (G Group[T]) Commutes[T any](a T, b T) bool {
	return G.Equals(G.Commutator(a, b), G.Identity())
}

func (G Group[T]) SubgroupGeneratedBy[T any](x T, max_iter int) []T, error {
	if G.Order() < 0 {
		return []T{}, fmt.Errorf("Error: Group order is negative.")
	}
	if max_iter <= 0 {
		return []T{}, fmt.Errorf("Max iterations is not positive")
	}
	iterations = 0
	elements := []T{G.Identity()}
	current = x
	for iterations <= max_iter {
		if current == G.Identity {
			break
		}
		elements = append(elements, current)
		current = G.Op(current, x)
	}
	return elements, nil
}


func (G Group[T]) IsFinite[T any]() bool {
	return G.Order() > 0
}
// Cayley tables?
//
// IsFinite: - switch (G.family) ;; if cyclic, easy
