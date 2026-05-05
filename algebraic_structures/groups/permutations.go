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


func isValidPermutation(x Permutation) bool {
	return isValidPermutationData(x.data)
}


func makePermutation(data []int) Permutation, error {
	if !isValidPermutationData(data) {
		return Permutation{}, fmt.Errorf("Invalid permutation.\nPlease ensure that all of 1,2,...,n appear in your list in a unique fashion.")
	}
	return Permutation{data: data}, nil
}


func permutationComposition(f Permutation, g Permutation, objects int) Permutation, error { 
	if len(f.data) != len(g.data) {
		// later, can add support for extending one perm to the other
		return Permutation{}, fmt.Errorf("Incompatible permutations")
	}
	if len(f.data) > objects {
		return Permutation{}, fmt.Errorf("Permutation acts on too many elements")
	}
	new_data := make([]int, len(f))
	for x, y := range f.data {
		new_data[x] = g.data[y]
	}
	return Permutation{data: new_data}, nil
}

func (p Permutation) invertPermutation() Permutation {
	inverted_data = make([]int, len(p.data))
	for x, y := range p.data {
		inverted_data[y] = x
	}
	return Permutation{data: inverted_data}
}

func permutationEqual(f Permutation, g Permutation) bool {
	// make this EqualTo method??
	return slices.Equal(f.data, g.data)
}


func (p Permutation) Sgn() int, error {
	inversions := 0
  for i := 0; i < len(p.data) - 1; i++ {
		if p.data[i+1] < p.data[i] {
			inversions += 1
		}
  }
	if inversions % 2 == 0 {
		return 1
	}
	return -1
}

// add error handling for invalid permutations. this may reduce speed. can potentially add in a "check_validity = True"



func (p Permutation) extend(n int) Permutation, error {
	if len(p.data) > n {
		return Permutation{}, fmt.Errorf("Cannot extend a permutation of length %d to one of length %d", len(p.data), n)
	}
	for i := len(p.data); i < n; i++ {

	}
}

// TODO:make sure that things that should be methods are indeed methods
