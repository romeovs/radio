package main

import (
	"fmt"
	"io"
	"os"
)

var w = 16

func main() {
	p := make([]byte, w)

	for {
		n, err := os.Stdin.Read(p)

		if err == io.EOF {
			os.Exit(0)
		}

		if err != nil {
			fmt.Printf("hex: error reading from stdin: %s\n", err)
			os.Exit(1)
		}

		for _, b := range p[:n] {
			fmt.Printf("%02x ", b)
		}

		for i := 0; i < w-n; i++ {
			fmt.Printf("   ")
		}

		fmt.Printf("     ")

		for _, b := range p[:n] {
			prchr(b)
		}

		for i := 0; i < w-n; i++ {
			fmt.Printf("  ")
		}

		fmt.Printf("\n")
	}
}

func prchr(b byte) {
	if b == 0x20 {
		fmt.Printf("␣")
		return
	}

	if b < 31 || b > 122 {
		fmt.Printf(".")
		return
	}

	if b == 127 {
		fmt.Printf("↩")
		return
	}

	fmt.Printf("%s", string([]byte{b}))
}
