package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	timeout = 400 * time.Millisecond
)

func main() {
	num := 22
	if len(os.Args) == 2 {
		n, err := strconv.ParseUint(os.Args[1], 10, 64)
		if err != nil {
			fmt.Printf("Could not parse pin number \"%s\"\n", os.Args[1])
			os.Exit(1)
		}
		num = int(n)
	} else if len(os.Args) > 2 {
		fmt.Println("Expected one argument: the pin number")
		os.Exit(1)
	}

	if num > 24 {
		fmt.Println("Select a pin number between 0 and 22")
		os.Exit(1)
	}

	// Use MCU pin 22, corresponds to GPIO 3 on the pi
	pin := rpio.Pin(num)

	// Open and map memory to access GPIO, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Disable edge event detection
	defer pin.Detect(rpio.NoEdge)

	// Unmap GPIO memory when done
	defer rpio.Close()

	// Set up pin as input pin
	pin.Input()

	// Pull up the pin
	pin.PullUp()

	// Enable falling edge event detection
	pin.Detect(rpio.FallEdge)

	fmt.Println("press a button...")

	for {
		time.Sleep(timeout)

		// Check if event occurred
		if pin.EdgeDetected() {
			fmt.Println("button pressed")
		}
	}
}
