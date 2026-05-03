package groups

import (
	"fmt"
	"slices"
)


type Permutation struct {
	data []int
}

func isValidPermutationData(data []int) bool {
	n := len(data)
	set := make(map[int]struct{})

	for i := 1; i <= n; i++ {
		set[data[i]] = struct{}{}
	}

	for i := 1; i <= n; i++ {
		_, exists := set[i] // might cause bug bc of walrus operator
		if !exists{
			return false
		}
	}
	return true
}

func isValidPermutation(p Permutation) bool {
	return isValidPermutationData(p.data)
}

func makePermutation(data []int) Permutation, error {
	if !isValidPermutationData(data) {
		return Permutation{}, fmt.Errorf("Invalid permutation.\nPlease ensure that all of 1,2,...,n appear in your list in a unique fashion.")
	}
	return Permutation{data: data}, nil
}


func permutationComposition(f Permutation, g Permutation) Permutation, error {
	// assumes that permutations are valid
	if len(f.data) != len(g.data) or Set(f.data) {
		// later, can add support for extending one perm to the other
		return Permutation{}, fmt.Errorf("Incompatible permutations") 
	}
	new_data := make([]int, len(f))
	for x, y := range f.data {
		new_data[x] = g.data[y]
	}
	return Permutation{data: new_data}, nil
}

func invertPermutation(f Permutation) Permutation {
	inverted_data = make([]int, len(f.data))
	for x, y := range f.data {
		inverted_data[y] = x
	}
	return Permutation{data: inverted_data}
}

func permutationEqual(f Permutation, g Permutation) bool {
	return slices.Equal(f.data, g.data)
}


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
		indicator: isValidPermutation
		op: permutationComposition,
		identity: idPermutation,
		inverse: invertPermutation,
		equals: permutationEqual,
		order: n,
		family: "symmetric",
	}
}


