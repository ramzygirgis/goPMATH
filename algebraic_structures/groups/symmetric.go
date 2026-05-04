package groups

import (
	"fmt"
	"slices"
)

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
