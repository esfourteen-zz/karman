package main

import (
	"./mcp3008"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	defer uiShutdown()

	var err error
	err = mcp3008.Setup(mcp3008.CHAN_0, 1000000)
	if err != nil {
		fail(err.Error())
	}

	initDisplay()

	var (
		tolerance, oldVal, psi float64 = .05, 0, 0
		chan0                  uint16  = 0
	)
	for {
		chan0, err = mcp3008.ReadADC(0)
		if err != nil {
			continue
		}
		psi = float64(toPsi(chan0))
		//fmt.Printf("D=%d\tA=%f\tPSI=%f\n", chan0, dToA(chan0), psi)
		if math.Abs(psi-oldVal) > tolerance {
			drawGauge(float32(psi))
			oldVal = psi
		}

		time.Sleep(30 * time.Millisecond)
	}
}

func barometricCalibration() float32 {
	return 100.53034
}

func toPsi(d uint16) float32 {
	kpa := float32(d) / 10.23 * 3.155
	psi := (kpa - barometricCalibration()) * 0.145037738
	return psi
}

func fail(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
