package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func parseArgs() (string, []string) {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	var command string
	var args []string
	str := string(stdin)
	if len(os.Args[1:]) >= 1 {
		command = os.Args[1]
		args = os.Args[2:]
	} else {
		command = ""
	}
	for _, arg := range strings.Split(str, "\n") {
		if arg != "" {
			args = append(args, arg)
		}
	}

	return command, args
}

func main() {
	command, args := parseArgs()

	if command != "" {
		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Println("ARGS: ", args)
	}
}
