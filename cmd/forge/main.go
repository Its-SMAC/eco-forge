package main

import "forge/internal/cli"

func main() {
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
