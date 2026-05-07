package groups

import (
	"fmt"
	"slices"
)


type SymmetricGroup struct {
	order int // order here is the n in S_n, rather than n!
}
// permutations should be created via NewPermutation to ensure validity


func NewSymmetricGroup(n int) SymmetricGroup, error{ 
	if n <= 0 {
		return SymmetricGroup{}, fmt.Errorf("Error: Cannot create a symmetric group on a %d elements", n)
	}
	return SymmetricGroup{order: n}, nil
}


func (G SymmetricGroup) IsValidPermutation(p permutation) error {
	// TODO: check that this catches all possible errors in all operations
	n := G.order
	if !isValidPermutation(data) {
		return fmt.Errorf("Error: Invalid permutation. \nPlease ensure that the data field contains all of the integers 1,2,...,n exactly once.")
	}
	if n < len(data) {
		return fmt.Errorf("Error: Permutation is not in S_n for n = %d.", n)
	}
	return nil
}


func (G SymmetricGroup) NewPermutation(data []int) permutation, error {
	if err := G.IsValidPermutation(p); err != nil {
		return permutation{}, err
	}
	n := G.order
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
	n := G.order
	p, _ := permutationComposition(a, b, n)
	return p
}


func (G SymmetricGroup) Inverse(p permutation) permutation {
	// ignore error, since we impose that permutations must be created using the NewPermutation method
	q = p.invertPermutation()
	q, _ := q.extend(G.order)
	return q
}


func (G SymmetricGroup) Contains(p permutation) bool {
	isValid := (G.IsValidPermutation(p) == nil)
	return isValid && len(p) == G.order
}


func (G SymmetricGroup) Identity() {
	data = make([]int, G.order)
	for i := 0; i < G.order; i++ {
		data[i] = i
	}
	return permutation{data: data}
}


func SymmetricGroup(n int) GenericGroup[permutation], error {
	if n <= 0 {
		return GenericGroup{}, fmt.Errorf("lol wdym you want symmetric group on %d elements", n)
	}
	idPermutationData := make([]int, n)
	for i := 0; i < n; i++ {
		idPermutationData[i] = i
	}
	idPermutation := permutation{idPermutationData}

	S_n := GenericGroup[permutation]{
		// indicator: symmetricGroupIndicator,
		op: permutationComposition,
		identity: idPermutation,
		inverse: invertPermutation,
		equals: permutationEqual,
		order: n,
		family: "symmetric",
	}
}

func (G GenericGroup[permutation]) Contains(p permutation) {
	if !p.isValidPermutation() {
		return false
	}
	return len(p.(permutation).data) <= G.order
}


func InclusionMap(G1 SymmetricGroup, G2 SymmetricGroup) (func(permutation) (permutation, error), error) {
	if G1.order > G2.order {
		return nil, fmt.Errorf("There is no inclusion mapping from the first group you passed to the second group you passed")
	}
	return func(p permutation) (permutation, error) {
		q, err := p.extend(G.order)
		return q, err
	}
}
