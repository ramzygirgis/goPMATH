package groups

import (
	"fmt"
	"slices"
)


type SymmetricGroup struct {
	degree int
}
// permutations should be created via NewPermutation to ensure validity


func NewSymmetricGroup(n int) SymmetricGroup, error{ 
	if n <= 0 {
		return SymmetricGroup{}, fmt.Errorf("Error: Cannot create a symmetric group on a %d elements", n)
	}
	return SymmetricGroup{degree: n}, nil
}


func (G SymmetricGroup) IsValidPermutation(p permutation) error {
	// TODO: check that this catches all possible errors in all operations
	n := G.degree
	if !isValidPermutation(data) {
		return fmt.Errorf("Error: Invalid permutation. \nPlease ensure that the data field contains all of the integers 1,2,...,n exactly once.")
	}
	if n < len(data) {
		return fmt.Errorf("Error: Permutation is not in S_n for n = %d.", n)}
	return nil
}


func (G SymmetricGroup) NewPermutation(data []int) permutation, error {
	if err := G.IsValidPermutation(p); err != nil {
		return permutation{}, err
	}
	n := G.degree
	p := permutation{data: data}
	if len(data) < n {
		var err error
		p, err = p.extend(n)
		if err != nil {
			return permutation{}, err
		}
	}
	return p, nil
}


func (G SymmetricGroup) Op(a, b permutation) permutation {
	// ignore error, since we impose that permutations must be created using the NewPermutation method
	n := G.degree
	p, _ := permutationComposition(a, b, n)
	return p
}


func (G SymmetricGroup) Identity() int {
	data = make([]int, G.degree)
	for i := 0; i < G.degree; i++ {
		data[i] = i
	}
	return permutation{data: data}
}


func (G SymmetricGroup) Inverse(p permutation) permutation {
	// ignore error, since we impose that permutations must be created using the NewPermutation method
	q = p.invertPermutation()
	q, _ := q.extend(G.degree)
	return q
}


func (G SymmetricGroup) Equals(f permutation, g permutation) bool {
	return permutationEqual(f, g)	
}


func (G SymmetricGroup) Order() int {
	n := G.degree
	if n > 20 {
		return -1
	}
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}


func (G SymmetricGroup) Family() string {
	return "symmetric"
}


func (G SymmetricGroup) Contains(p permutation) bool {
	isValid := (G.IsValidPermutation(p) == nil)
	return isValid && len(p) == G.degree
}


func (G SymmetricGroup) Degree() {
	return G.degree
}


func InclusionMap(G1 SymmetricGroup, G2 SymmetricGroup) (func(permutation) (permutation, error), error) {
	if G1.degree > G2.degree {
		return nil, fmt.Errorf("There is no inclusion mapping from the first group you passed to the second group you passed")
	}
	return func(p permutation) (permutation, error) {
		q, err := p.extend(G.degree)
		return q, err
	}
}
