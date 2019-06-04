package main

import (
	"fmt"
	"github.com/xyproto/swish"
	"os"
	"strings"
	"text/tabwriter"
)

func abs(x float64) float64 {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	fmt.Println("Optimized Swish precision")
	fmt.Println()

	w := tabwriter.NewWriter(os.Stdout, 15, 0, 0, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "Swish\tSwishPrecise\tDifference\t")
	var sumDiff float64
	var minDiff float64
	var maxDiff float64
	counter := 0
	for x := -5.0; x <= 3.0; x += 0.1 {
		a := swish.Swish(x)
		b := swish.SwishPrecise(x)
		diff := abs(b - a)
		fmt.Fprintf(w, "%.4f\t%.4f\t%.4f\t\n", a, b, diff)
		sumDiff += diff
		counter++
		if diff > maxDiff {
			maxDiff = diff
		}
		if diff < minDiff || minDiff == 0.0 {
			minDiff = diff
		}
	}
	w.Flush()

	avgDiff := sumDiff / float64(counter)

	fmt.Println()
	fmt.Println(strings.Repeat("-", 21))
	fmt.Printf("Minimum error: %3.4f\n", minDiff)
	fmt.Printf("Average error: %3.4f\n", avgDiff)
	fmt.Printf("Maximum error: %3.4f\n", maxDiff)
}
