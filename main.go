package main

import (
	"fmt"
	"github.com/khmarbaise/blinkt/gpio"
)

func main() {
	fmt.Println("Main of package gpio")
	rasberry3Plus := gpio.Rasberry3Plus{}
	exported := gpio.IsGpioPinExported(rasberry3Plus, 23)
	fmt.Printf("exported: %v\n", exported)
}
