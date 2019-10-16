package exec

import (
	"fmt"
	"io"
	"os/exec"
)

// Command runs the command.
// Returns a Reader that contains the stdout of the command.
func Command(name string, args ...string) (*Cmd, error) {
	return Pipe(nil, name, args...)
}

// Pipe creates and runs a command and pipes s into its stdin.
// Returns a Reader that contains the stdout of the command.
func Pipe(s io.Reader, name string, args ...string) (*Cmd, error) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if s != nil {
		cmd.Stdin = s
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return &Cmd{
		stdout: stdout,
		cmd:    cmd,
	}, nil
}

// Cmd wraps exec.Cmd
type Cmd struct {
	stdout io.ReadCloser
	cmd    *exec.Cmd
}

// Read data from the Cmds stdout.
// If EOF is reached, the underlying command will be released.
func (cmd *Cmd) Read(p []byte) (int, error) {
	n, err := cmd.stdout.Read(p)

	if err != nil {
		fmt.Println("STOPPING")
		cmd.cmd.Wait()
		return 0, err
	}

	fmt.Printf("read %s bytes\n", n)
	return n, err
}

// Close the underlying command.
func (cmd *Cmd) Close() error {
	return cmd.cmd.Wait()
}
