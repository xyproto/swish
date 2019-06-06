package swish

// Thanks https://codingforspeed.com/using-faster-exponential-approximation/
func exp256(x float64) float64 {
	x2 := x
	x2 /= 256.0
	x2 += 1.0
	x2 *= x2
	x2 *= x2
	x2 *= x2
	x2 *= x2
	x2 *= x2
	x2 *= x2
	x2 *= x2
	x2 *= x2
	return x2
}

// Sigmoid is the 1 / (1 + exp(-x)) activation function
func Sigmoid(x float64) float64 {
	// Uses exp256 instead of math.Exp
	return 1.0 / (1.0 + exp256(-x))
}
