package main

import (
	"fmt"
	"forge/internal/cli"
	"os"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
