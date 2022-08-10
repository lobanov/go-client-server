package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Verbose bool `short:"v" long:"verbose" description:"Verbose output"`
	Port    uint `short:"p" long:"port" description:"Port to listen to" default:"1234"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	_, err := parser.Parse()
	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)

		default:
			os.Exit(1)
		}
	}
	fmt.Println(options)
}
