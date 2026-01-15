package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// getPrompt returns the shell prompt string
func getPrompt() string {
	return "shell>"
}

// executeCommand executes an external command and returns its output
func executeCommand(cmdStr string) (string, error) {
	cmdStr = strings.TrimSpace(cmdStr)
	if cmdStr == "" {
		return "", nil
	}

	// Parse command and arguments
	parts := strings.Fields(cmdStr)
	if len(parts) == 0 {
		return "", nil
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// runShellWithIO runs the shell with custom input/output for testing
func runShellWithIO(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	
	for {
		// Display prompt
		fmt.Fprint(output, getPrompt(), " ")
		
		// Read user input
		if !scanner.Scan() {
			break
		}
		
		line := scanner.Text()
		line = strings.TrimSpace(line)
		
		// Check for exit command
		if line == "exit" {
			break
		}
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Execute command and display output
		result, err := executeCommand(line)
		if err != nil {
			fmt.Fprintf(output, "Error: %v\n", err)
		}
		if result != "" {
			fmt.Fprint(output, result)
		}
	}
}

// runShell runs the interactive shell
func runShell() {
	runShellWithIO(os.Stdin, os.Stdout)
}

func main() {
	runShell()
}
