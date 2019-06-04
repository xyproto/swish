package swish

import (
	"math"
)

// The SWISH activation function
func F(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
