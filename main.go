package main

import (
	"os"

	"github.com/DongLieu/golang-blockchain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
