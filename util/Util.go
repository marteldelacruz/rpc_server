package util

import (
	"bufio"
	"os"
)

const (
	PROTOCOL = "tcp"
	PORT     = ":9999"
)

type Args struct {
	Name    string
	Subject string
	Grade   float64
}

// Scans from the user input and then returns the
// complete string
func ScanString() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	return scanner.Text()
}
