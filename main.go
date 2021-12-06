package main

import (
	"fmt"
	"os"

	"github.com/goodmustache/appli.ed/cmd"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	command := &cmd.Execute{}
	parser := flags.NewParser(command, flags.HelpFlag)

	extraArgs, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = command.Execute(extraArgs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "appli.ed error", err.Error())
		os.Exit(1)
	}
}
