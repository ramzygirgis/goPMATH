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

func Commutator[T any](G Group[T], a T, b T) T {
	ab := G.Op(a, b)
	ba := G.Op(b, a)
	ba_inverse = G.Inverse(ba)
	return G.Op(ab, ba_inverse)
}

func Commutes[T any](G Group[T], a T, b T) bool {
	return G.Equals(Commutator(G, a, b), G.Identity())
}

func (G AbstractGroup[T]) SubgroupGeneratedBy[T any](x T, max_iter int) Group[T] {
	
	if G.order != 0 {
		accum = 
		o := 1
		while 
	}
	if max_iter <= 0 {

	}
}


func (G IsFinite() bool {

}
// Cayley tables?
//
// IsFinite: - switch (G.family) ;; if cyclic, easy
