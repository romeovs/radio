package gpio

import "github.com/stianeikeland/go-rpio"

var (
	rotaryPinA = 1
	rotaryPinB = 2
	rotaryPinC = 3
	rotaryPinD = 4
)

// rotarySwitch decodes the state of the rotary switch.
//
// See https://docs-emea.rs-online.com/webdocs/0ff6/0900766b80ff6d94.pdf
//
// In the version of the coded rotary switch I am using (Lorlin BCK1002), the code is given
// in hexadecimal binary compliment.
//
// This it the truth table from that document:
//
//          TERMINALS
//  VALUE   ABCD
//     1    1000
//     2    0100
//     3    1100
//     4    0010
//     5    1010
//     6    0110
//     7    1110
//     8    0001
//     9    1001
//    10    0101
//    11    1101
//    12    0011
//    13    1011
//    14    0111
//    15    1111
//    16    0000
func rotarySwitch(A, B, C, D rpio.State) int {
	dail := 0b0000

	if A == rpio.High {
		dail = dail | 0b0001
	}

	if B == rpio.High {
		dail = dail | 0b0010
	}

	if C == rpio.High {
		dail = dail | 0b0100
	}

	if D == rpio.High {
		dail = dail | 0b1000
	}

	if dail == 0 {
		return 16
	}

	return dail
}
