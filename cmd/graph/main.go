package main

import (
	tm "github.com/buger/goterm"
	"github.com/xyproto/swish"
)

func main() {
	tm.Clear()
	tm.MoveCursor(0, 0)

	tm.Println("\n" + tm.Bold(tm.Color("The Sigmoid activation function: ", tm.WHITE)) + tm.Color("1 / (1 + exp(-x))", tm.BLUE) + "\n")

	// Use 70% of the terminal width
	width := int(float64(tm.Width()) * 0.7)
	height := int(float64(width) * 0.3)

	chart := tm.NewLineChart(width, height)
	chart.Flags = tm.DRAW_RELATIVE
	data := new(tm.DataTable)
	data.AddColumn(" x ")
	data.AddColumn(" y ")

	for x := -5.0; x <= 3.0; x += (8.0 / float64(width)) {
		data.AddRow(x, swish.Sigmoid(x))
	}

	tm.Println(chart.Draw(data))
	tm.Flush()

	tm.Println("\n" + tm.Bold(tm.Color("The Swish activation function: ", tm.WHITE)) + tm.Color("x / (1 + exp(-x))", tm.BLUE) + "\n")

	chart = tm.NewLineChart(width, height)
	chart.Flags = tm.DRAW_RELATIVE
	data = new(tm.DataTable)
	data.AddColumn(" x ")
	data.AddColumn(" y ")

	for x := -5.0; x <= 3.0; x += (8.0 / float64(width)) {
		data.AddRow(x, swish.F(x))
	}

	tm.Println(chart.Draw(data))
	tm.Flush()

}
