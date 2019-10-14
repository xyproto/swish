//+build amd64

package swish

// go: noescape
// SwishAssembly is the swish function, written in hand-optimized assembly
func SwishAssembly(x float64) float64
