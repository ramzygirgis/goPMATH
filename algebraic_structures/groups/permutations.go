package groups

import (
	"fmt"
	"slices"
)


type permutation struct {
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


func permutationComposition(f permutation, g permutation, objects int) permutation, error { 
	if len(f.data) > objects or len(g.data) > objects {
		return permutation{}, fmt.Errorf("Permutation acts on too many elements")
	}
	if len(f.data) <= len(g.data) {
		f, err = f.extend(len(g.data))
		if err != nil {
			return permutation{}, err
		}
	} else if len(g.data) < len(f.data) {
		g, err = g.extend(len(f.data))
		if err != nil {
			return permutation{}, err
		}
	}
	new_data := make([]int, len(f))
	for x, y := range f.data {
		new_data[x] = g.data[y]
	}
	return permutation{data: new_data}, nil
}

func (p permutation) invertPermutation() permutation {
	inverted_data = make([]int, len(p.data))
	for x, y := range p.data {
		inverted_data[y] = x
	}
	return permutation{data: inverted_data}
}


func permutationEqual(f permutation, g permutation) bool {
	// make this EqualTo method??
	return slices.Equal(f.data, g.data)
}


func (p permutation) Sgn() int, error {
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




func (p permutation) extend(n int) permutation, error {
	if len(p.data) > n {
		return permutation{}, fmt.Errorf("Cannot extend a permutation of length %d to one of length %d", len(p.data), n)
	}
	for i := len(p.data); i < n; i++ {
		p.data[i] = i
	}
	return p, nil
}

// TODO:make sure that things that should be methods are indeed methods
//
// TODO: update certain methods to pass by ref
