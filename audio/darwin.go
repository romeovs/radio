// +build darwin

package audio

var (
	cmd  = "paplay"
	args = []string{
		"--raw",
		"--fix-rate",
		"--latency-msec=10",
	}
)
