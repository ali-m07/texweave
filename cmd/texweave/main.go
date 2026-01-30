package main

import (
	"os"

	"github.com/ali-m07/texweave/cmd/texweave/root"
)

func main() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
