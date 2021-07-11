package main

import (
	"flag"
	"github.com/lroman242/envdir/utils"
	"log"
	"os"
)

func main()  {
	flag.Parse()

	args := flag.Args()
	log.Printf("%+v\n", args)

	if len(args) < 2 {
		log.Fatalf("invalid number of arguments. expect at least 2 arguments")
	}

	env, err := utils.ReadDir(args[0])
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", env)
	log.Printf("%+v\n", args[1:])

	os.Exit(utils.RunCommand(args[1:], env))
}
