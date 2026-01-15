package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// parseInput splits the input string into command and arguments
func parseInput(input string) (string, []string) {
	// Trim spaces and split by whitespace
	input = strings.TrimSpace(input)
	if input == "" {
		return "", []string{}
	}
	
	parts := strings.Fields(input)
	
	cmd := parts[0]
	args := []string{}
	if len(parts) > 1 {
		args = parts[1:]
	}
	
	return cmd, args
}

// executeCommand executes an external command and returns its output
func executeCommand(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("empty command")
	}
	
	cmd, args := parseInput(input)
	
	command := exec.Command(cmd, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr
	
	err := command.Run()
	if err != nil {
		return stderr.String(), err
	}
	
	return out.String(), nil
}

// runShell runs the interactive shell
func runShell() {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		// Display the prompt
		fmt.Print("shell> ")
		
		// Read user input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			continue
		}
		
		// Trim newline
		input = strings.TrimSpace(input)
		
		// Skip empty input
		if input == "" {
			continue
		}
		
		// Check for exit command
		if input == "exit" || input == "quit" {
			fmt.Println("Goodbye!")
			break
		}
		
		// Execute the command
		output, err := executeCommand(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			if output != "" {
				fmt.Fprint(os.Stderr, output)
			}
		} else {
			// Display the output
			fmt.Print(output)
		}
	}
}

func main() {
	runShell()
}
