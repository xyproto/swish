package swish

import (
	"fmt"
	"math"
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

func TestSwish(t *testing.T) {
	diff := math.Abs(Swish(0.7) - SwishPrecise(0.7))
	fmt.Printf("%3.5f\n", diff)
	if diff >= 0.0002 {
		t.Fail()
	}
	diff = math.Abs(Swish(0.3) - SwishPrecise(0.3))
	fmt.Printf("%3.5f\n", diff)
	if diff >= 0.00002 {
		t.Fail()
	}
}

func TestGaussian(t *testing.T) {
	// Check the accuracy, compared to the result from this
	// numpy expression (where x = 2.0):
	//
	// np.exp(-np.multiply(x, x) / 2.0)

	diff := math.Abs(Gaussian01(2.0) - 0.1353352832366127)
	if diff >= 0.002 {
		t.Fail()
	}
}
