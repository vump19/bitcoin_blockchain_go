package main

import (
	"fmt"
	"os"
)

type CLI struct {
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}
func (cli *CLI) Run() {
	cli.validateArgs()

	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Println("NODE_ID env. var is not set!")
		os.Exit(1)
	}

	//getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

}
