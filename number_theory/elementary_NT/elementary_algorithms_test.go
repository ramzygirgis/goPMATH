package elementaryNT

import "testing"

func FuzzAbs(t *testing.F) {
	f.Add(0)
	f.Add(-1)
	f.Add(1)
	f.Add(10000)
	f.Add(-111000)

	f.Fuzz(func(t *testing.T, a int) {
		res := abs(a)
		if res < 0 {
			t.Fatalf("expected positive result for abs(%d), got %d", a, res)
		}
	})
}

func FuzzGcd(f *testing.F) {
    f.Add(12, 18)
		f.Add(13, 1)
		f.Add(1,0)
		f.Add(3, -18)
		f.Add(189373082, 1282109)
		f.Add(-19298,-290)
		f.Add(0,0)

    f.Fuzz(func(t *testing.T, a, b int) {
			g := Gcd(a, b)
			
			if g == 0 && (a !=0 || b != 0) {
				t.Fatalf("gcd(%d, %d) is apparently zero?", a, b)
			}
			if g != 0 && (a % g != 0 || b % g != 0) {
				t.Fatalf("gcd(%d, %d)=%d doesn't divide one of %d or %d", a, b, g, a, b)
			}
			if g < 0 {
				t.Fatalf("gcd(%d, %d)=%d appears to be negative...", a, b, g)
			}
    })
}
