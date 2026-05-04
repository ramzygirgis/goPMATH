package groups

import "fmt"

func cyclicAdd(a, b int) int {
	return a + b
}

func cyclicInverse(a int) int {
	return -a
}

func CyclicGroup(n int) (GenericGroup[int], error) {
	if n < 0 {
		return GenericGroup[int]{}, fmt.Errorf("lol wdym you want a group of order < 0?")
	}

	var equals func(int, int) bool
	if n == 0 {
		equals = func(a, b int) bool {
			return a == b
		}
	} else {
		equals = func(a, b int) bool {
			return a % n == b % n
		}
	}

	C_n := GenericGroup[int]{
		op: cyclicAdd,
		identity: 0,
		inverse: cyclicInverse,
		equals: equals,
		order: n,
		family: "cyclic",
		contains: func(x int) bool { // useless here tbh
			return True
		}
	}

func (G GenericGroup[int]) Contains(x int) bool {
	return G.contains(x)
}

func (G GenericGroup[int]) Equals(x, y int) bool {
	return G.equals(x, y)
}


func GetAdditiveOrder(x, n int) int {
	// gets the additive order of x mod n
	o := 1
	accum := x
	for accum % n != 0 {
		accum += x
		o += 1
	}
	return o
}
