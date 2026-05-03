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
	return func(a, b int) bool {
		return a % n == b % n
	}
}

func isInt(v any) bool {
	// 5.0 will not be considered an integer; always convert floats.
    switch v.(type) {
    case int, int8, int16, int32, int64:
        return true
    default:
        return false
    }
}


func CyclicGroup(n int) GenericGroup[int], error {
	if n < 0 {
		return GenericGroup{}, fmt.Errorf("lol wdym you want a group of order < 0?")
	}
	C_n := GenericGroup[int]{
		indicator: isInt
		op: cyclicAdd,
		identity: 0,
		inverse: cyclicInverse,
		equals: createCyclicEquals(n),
		order: n,
		family: "cyclic"
	}
	return C_n, nil
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
