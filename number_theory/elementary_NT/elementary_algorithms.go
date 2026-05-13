package elementaryNT


import "fmt"


func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}


func Gcd(a, b int) int {
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


func Lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return abs(a / Gcd(a, b) * b)
}


func pAdicVal(n int, p int) (int, error) {
	if n == 0 {
		return 0, fmt.Errorf("p-adic valuation of 0 is undefined") // band aid fix for now
	}
	if !isPrime(p) { // implement this
		return 0, fmt.Errorf("%d is not prime", p)
	}
	prod := abs(n)
	v_p := 0
	for prod % p == 0 {
		prod = prod / p
		v_p += 1
	}
	return v_p, nil
}


func pAdicValRational(a, b, p int) (int, error) {
	if !isPrime(p){
		return 0, fmt.Errorf("%d is not prime", p)
	}
	x_a, err := pAdicVal(a, p)
	if err != nil {
		return 0, err
	}
	x_b, err := pAdicVal(b, p)
	if err != nil {
		return 0, err
	}
	return x_a - x_b, nil
}


func pAdicValFactorial(n int, p int) (int, error) {
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
  return ans, nil
}


func pAdicValBinomial(n, k, p int) (int, error) {
	if !isPrime(p) {
		return 0, fmt.Errorf("%d is not prime", p)
	}
	if n < k { // HANDLE NEGATIVES
		return pAdicVal(0, p)
	}
	x_1, err := pAdicValFactorial(n, p)
	if err != nil {
		return 0, err
	}
	x_2, err := pAdicValFactorial(k, p)
	if err != nil {
		return 0, err
	}
	x_3, err := pAdicValFactorial(n - k, p)
	if err != nil {
		return 0, err
	}
	return a - b - c, nil
}
