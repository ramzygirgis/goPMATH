package groups

import "fmt"

type CyclicGroup struct { // this will satisfy the interface for Group[int]. we can make a matrix version, and implement isomorphisms
	order int // for \mathbb{Z}, order = 0. TODO order shouldn't be negative; find a way to workaround
}


func CreateCyclicGroup(n int) CyclicGroup, error { 
	if n < 0 {
		return CyclicGroup{}, fmt.Errorf("Hello??? Cyclic group of negative order???") 
	}
	return CyclicGroup{order: n}, nil
}


func (G CyclicGroup) Op(a, b int) int {
	return (a + b) % G.order
}


func (G CyclicGroup) Identity() int {
	return 0
}


func (G CyclicGroup) Inverse(a int) int {
	n := G.order
	if n == 0 {
			return -a
		}
	return ((-a % n) + n) % n
}


func (G CyclicGroup) Equals(a, b int) bool {
	n := G.order
	if n == 0 {
		return a == b
	}
	return a % n == b % n
}


func (G CyclicGroup) Order() int {
	return G.order
}


func (G CyclicGroup) Family() string {
	return "cyclic"
}


func (G CyclicGroup) Contains(a int) {
	return true
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
