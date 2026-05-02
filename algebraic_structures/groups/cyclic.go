package groups

import "fmt"

func cyclicAdd(a, b int) int {
	return a + b
}

func cyclicInverse(a int) int {
	return -a
}

func createCyclicEquals(n int) {
	if n == 0 {
		return func(a, b int) int {
			return a == b
		}
	}
	return func(a, b int) int {
		return a % n == b % n
	}
}

// maybe, can do inheritance, where this type of group gains an extra field; order


func CyclicGroup(n int) GenericGroup[int], error {
	if n < 0 {
		return GenericGruop{}, fmt.Errorf("lol wdym you want a group of order < 0?")
	}
	C_n := GenericGroup[int]{
		op: cyclicAdd,
		identity: 0,
		inverse: cyclicInverse,
		equals: createCyclicEquals(n),
		order: n,
		family: "cyclic"
	}
	return C_n, nil
}

func GetAdditiveOrder(x, n int) o int {
	// gets the additive order of x mod n
	o := 1
	accum := x
	while accum % n != 0 {
		accum += x
		o += 1
	}
	return o
}
