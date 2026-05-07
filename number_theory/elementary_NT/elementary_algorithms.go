package elementaryNT


import "fmt"


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


func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a * b / gcd(a, b)
}


func pAdicVal(n int, p int) int, error {
	if !isPrime(p) { // implement this
		return 0, fmt.Errorf("%d is not prime", p)
	}
	prod = n
	v_p = 0
	while prod % p == 0 {
		prod = prod / p
		v_p += 1
	}
	return v_p, nil
}


func pAdicValFactorial(n int, p int) {
	// returns the p-adic valuation of n!
	if !isPrime(p) {
		return 0, fmt.Errorf("%d is not prime", p)
	}
	if n < 0 {
		return 0, fmt.Errorf("%d is negative, we not tryna use gamma to compute n!", n)
	}
  ans := 0
  for n > 0 {
       n /= p
      ans += n
  }
  return ans
}
