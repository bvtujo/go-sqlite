package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var validCommands = map[string]func(){
	".exit": handleExitCmd,
}

func handleExitCmd() {
	fmt.Println("Bye!")
	os.Exit(0)
}

func printPrompt() { fmt.Printf("db > ") }

func flush() { fmt.Println("") }

func getInput(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')
	out := strings.TrimSpace(input)
	return out
}

func parseCommand(cmd string) (func(), error) {
	cmdFunc, ok := validCommands[cmd]
	if !ok {
		return nil, fmt.Errorf("Unrecognized command %s", cmd)
	}
	return cmdFunc, nil
}

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		printPrompt()
		input := getInput(r)
		cmd, err := parseCommand(input)
		if err != nil {
			fmt.Printf("%s", err.Error())
			flush()
			continue
		}
		cmd()
		flush()
	}
}
