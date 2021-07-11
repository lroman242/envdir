package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lroman242/envdir/utils"
)

const minNumberOfExpectedArguments = 2

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "envdir tool will parse environment variables from directory and run provided"+
			" command with arguments and parsed environment variables")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s /path/to/directory/ command [options|args...] \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	args := flag.Args()

	if len(args) < minNumberOfExpectedArguments {
		log.Fatalf("invalid number of arguments. expect at least 2 arguments")
	}

	env, err := utils.ReadDir(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(utils.RunCommand(args[1:], env))
}
