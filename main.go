package main

import (
	"os"

	"github.com/tensor-programming/blockchain-golang/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
