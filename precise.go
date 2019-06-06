package swish

import "math"

// SigmoidPrecise is the 1 / (1 + exp(-x)) activation function, using math.Exp
func SigmoidPrecise(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// SwishPrecise is the x / (1 + exp(-x)) activation function, using math.Exp
func SwishPrecise(x float64) float64 {
	return x / (1.0 + math.Exp(-x))
}
