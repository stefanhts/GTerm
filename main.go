package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Input loop
func main() {
	index := 0
	var commands []string
	//TODO: add command history
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := executeInput(input); err != nil {
			index += 1
			commands = append(commands, input)
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// Run the input command
func executeInput(input string) error {
	command := strings.TrimSuffix(input, "\n")
	args := strings.Split(command, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("cd: not enough arguments")
		}
		return os.Chdir(args[1])
	case "cow":
		var wisdom string
		if len(args) < 2 {
			wisdom = ""
		} else {
			wisdom = strings.Join(args[1:], " ")
		}
		return cow(wisdom)
	case "!!":
		fmt.Print("\033[H\033[2J")
		return nil
	case "exit":
		os.Exit(0)
	}

	result := exec.Command(args[0], args[1:]...)

	result.Stderr = os.Stderr
	result.Stdout = os.Stdout

	return result.Run()
}

//Room for custom commands
func cow(wisdom string) error {
	if len(wisdom) < 1 {
		wisdom = "I don't know how to handle this"
	}
	first := " "
	second := " "
	i := 0
	for i < len(wisdom)+2 {
		first += "_"
		second += "-"
		i = i + 1
	}

	fmt.Println(first)
	fmt.Println("< " + wisdom + " >")
	fmt.Println(second)
	fmt.Println("        \\   ^__^")
	fmt.Println("         \\  (oo)\\_______")
	fmt.Println("            (__)\\       )\\/\\")
	fmt.Println("                ||----w |")
	fmt.Println("                ||     ||")
	return nil
}
