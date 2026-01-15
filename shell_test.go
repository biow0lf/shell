package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrompt(t *testing.T) {
	expected := "shell>"
	actual := getPrompt()
	if actual != expected {
		t.Errorf("Expected prompt '%s', got '%s'", expected, actual)
	}
}

func TestExecuteCommand(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		wantErr  bool
		contains string
	}{
		{
			name:     "echo command",
			command:  "echo hello",
			wantErr:  false,
			contains: "hello",
		},
		{
			name:    "invalid command",
			command: "nonexistentcommand12345",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := executeCommand(tt.command)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if tt.contains != "" && !strings.Contains(output, tt.contains) {
					t.Errorf("Expected output to contain '%s', got '%s'", tt.contains, output)
				}
			}
		})
	}
}

func TestRunShell(t *testing.T) {
	input := "echo test\nexit\n"
	inputReader := strings.NewReader(input)
	var outputBuffer bytes.Buffer

	runShellWithIO(inputReader, &outputBuffer)

	output := outputBuffer.String()
	
	// Check that prompt is displayed
	if !strings.Contains(output, "shell>") {
		t.Errorf("Expected output to contain prompt 'shell>', got: %s", output)
	}
	
	// Check that command output is displayed
	if !strings.Contains(output, "test") {
		t.Errorf("Expected output to contain 'test', got: %s", output)
	}
}
