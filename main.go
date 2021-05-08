package main

import (
	"fmt"
	"github.com/khmarbaise/blinkt/gpio"
)

const (
	// DAT is the Data pin for Blinkt
	DAT int = 23

	// CLK is the clock pin for Blinkt
	CLK int = 24
)

func main() {
	fmt.Println("Main of package gpio")
	rasberry3Plus := gpio.Rasberry3Plus{}
	gpio.PinMode(rasberry3Plus, DAT)
	gpio.PinMode(rasberry3Plus, CLK)
	fmt.Println("done.")
}
