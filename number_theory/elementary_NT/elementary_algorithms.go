package elementaryNT


func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// AKS, carmicheal, lucas lehmer


func gcd(a, b int) int {
	a = abs(a)
	b = abs(b)

	if b == 0 {
		return a
	}
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
