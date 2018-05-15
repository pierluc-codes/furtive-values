package main

import (
	"fmt"
	"os"

	"github.com/plcstpierre/furtive-values/internal/text"
)

func main() {
	returnCode := 0

	defer func() {
		os.Exit(returnCode)
	}()

	err := text.ProcessStream(os.Stdin, os.Stdout)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		returnCode = -1
	}
}
