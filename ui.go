package main

import (
	"fmt"
	"github.com/ajstarks/openvg"
)

var (
	width, height int
	w2, h2, w     openvg.VGfloat
)

func initDisplay() {
	width, height = openvg.Init()
	w2 = openvg.VGfloat(width / 2)
	h2 = openvg.VGfloat(height / 2)
	w = openvg.VGfloat(width)
}

func drawGauge(val float32) {
	openvg.Start(width, height)     // Start the picture
	openvg.BackgroundColor("black") // Black background
	openvg.FillRGB(44, 100, 232, 1) // Big blue marble
	openvg.Circle(w2, 0, w)         // The "world"
	openvg.FillColor("white")       // White text
	openvg.TextMid(w2, h2, fmt.Sprintf("%2.2f", val), "serif", width/10)
	openvg.End()
}

func uiShutdown() {
	openvg.Finish()
}
