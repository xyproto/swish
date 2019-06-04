package swish

import (
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
