package groups

import (
	"fmt"
	"slices"
)


type SymmetricGroup struct {
	order int // order here is the n in S_n, rather than n!
}
// permutations should be created via MakePermutation to ensure validity


func CreateSymmetricGroup(n int) SymmetricGroup, error{ 
	if n <= 0 {
		return SymmetricGroup{}, fmt.Errorf("Cannot create a symmetric group on a %d elements", n)
	}
	return SymmetricGroup{order: n}, nil
}


func (G SymmetricGroup) Op(a, b permutation) permutation, error {
	n := G.order
	p, err := permutationComposition(a, b, n)
	return p, err
}

func (G SymmetricGroup) Inverse(p permutation) permutation, error {
	q = p.invertPermutation()
	q, err := q.extend(G.order)
	if err != nil {
		return permutation{}, err
	}
	return q, err
}

func (G SymmetricGroup) Contains(p permutation) bool {
	return p.isValidPermutation() && len(p) == G.order
}



// TODO:continue implementing symmetric group, modify interface for groups to contain error output, modify cyclic to contain errors

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
	if !isValidPermutation(p) {
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
