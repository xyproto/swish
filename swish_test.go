package swish

import (
	"fmt"
	"testing"
)

// benchmark using exp256
func BenchmarkSwish07(b *testing.B) {
	// run the Swish function b.N times
	for n := 0; n < b.N; n++ {
		Swish(0.7)
	}
}

// benchmark using exp256
func BenchmarkSwish03(b *testing.B) {
	// run the Swish function b.N times
	for n := 0; n < b.N; n++ {
		Swish(0.3)
	}
}

// benchmark using math.Exp
func BenchmarkSwishPrecise07(b *testing.B) {
	// run the SwishPrecise function b.N times
	for n := 0; n < b.N; n++ {
		SwishPrecise(0.7)
	}
}

// benchmark using math.Exp
func BenchmarkSwishPrecise03(b *testing.B) {
	// run the SwishPrecise function b.N times
	for n := 0; n < b.N; n++ {
		SwishPrecise(0.3)
	}
}

func abs(a float64) float64 {
	if a >= 0 {
		return a
	}
	return -a
}

func TestSwish(t *testing.T) {
	diff := abs(Swish(0.7) - SwishPrecise(0.7))
	fmt.Printf("%3.5f\n", diff)
	if diff >= 0.0002 {
		t.Fail()
	}
	diff = abs(Swish(0.3) - SwishPrecise(0.3))
	fmt.Printf("%3.5f\n", diff)
	if diff >= 0.00002 {
		t.Fail()
	}
}
