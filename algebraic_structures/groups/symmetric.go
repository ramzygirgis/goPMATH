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


func (G SymmetricGroup) Op(a, b Permutation) Permutation, error {
	n := G.order
	p, err := permutationComposition(a, b, n)
	return p, err
}

func (G SymmetricGroup) Inverse(p Permutation) Permutation, error {
	q = p.invertPermutation()
	q, err := q.extend(G.order)
	if err != nil {
		return Permutation{}, err
	}
	return q, err
}



// TODO:continue implementing symmetric group, modify interface for groups to contain error output, modify cyclic to contain errors

func SymmetricGroup(n int) GenericGroup[Permutation], error {
	if n <= 0 {
		return GenericGroup{}, fmt.Errorf("lol wdym you want symmetric group on %d elements", n)
	}
	idPermutationData := make([]int, n)
	for i := 0; i < n; i++ {
		idPermutationData[i] = i
	}
	idPermutation := Permutation{idPermutationData}

	S_n := GenericGroup[Permutation]{
		// indicator: symmetricGroupIndicator,
		op: permutationComposition,
		identity: idPermutation,
		inverse: invertPermutation,
		equals: permutationEqual,
		order: n,
		family: "symmetric",
	}
}

func (G GenericGroup[Permutation]) Contains(p Permutation) {
	if !isValidPermutation(p) {
		return false
	}
	return len(p.(Permutation).data) <= G.order
}
