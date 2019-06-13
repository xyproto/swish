// +build amd64

package swish

import (
	"testing"
)

// benchmark using handwritten assembly
func BenchmarkSwishAssembly07(b *testing.B) {
	// run the SwishPrecise function b.N times
	for n := 0; n < b.N; n++ {
		SwishAssembly(0.7)
	}
}

// benchmark using handwritten assembly
func BenchmarkSwishAssembly03(b *testing.B) {
	// run the SwishPrecise function b.N times
	for n := 0; n < b.N; n++ {
		SwishAssembly(0.3)
	}
}
