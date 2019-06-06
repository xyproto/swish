//+build !amd64

package swish

// Swish is the x / (1 + exp(-x)) activation function, using math.Exp
// go: noescape
func Swish(x0 float64) float64 {
	x1 := x0
	x2 := x0
	x1 *= -0.00390625 // (1 / -256.0)
	x1 += 1.0
	x1 *= x1
	x1 *= x1
	x1 *= x1
	x1 *= x1
	x1 *= x1
	x1 *= x1
	x1 *= x1
	x1 *= x1
	x1 += 1.0
	x1 /= x2
	return x1
}
