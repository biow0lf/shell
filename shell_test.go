package main

import (
	"strings"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "simple echo command",
			input:   "echo hello",
			wantErr: false,
		},
		{
			name:    "ls command",
			input:   "ls",
			wantErr: false,
		},
		{
			name:    "empty command",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := executeCommand(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("executeCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && output == "" {
				t.Errorf("executeCommand() expected output but got empty string")
			}
		})
	}
}

func TestExecuteCommandOutput(t *testing.T) {
	output, err := executeCommand("echo test")
	if err != nil {
		t.Fatalf("executeCommand() unexpected error: %v", err)
	}
	
	if !strings.Contains(output, "test") {
		t.Errorf("executeCommand() output = %v, want to contain 'test'", output)
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantCmd string
		wantArgs []string
	}{
		{
			name:    "simple command",
			input:   "echo hello",
			wantCmd: "echo",
			wantArgs: []string{"hello"},
		},
		{
			name:    "command with multiple args",
			input:   "ls -la /tmp",
			wantCmd: "ls",
			wantArgs: []string{"-la", "/tmp"},
		},
		{
			name:    "command with no args",
			input:   "pwd",
			wantCmd: "pwd",
			wantArgs: []string{},
		},
		{
			name:    "command with extra spaces",
			input:   "  echo   hello  ",
			wantCmd: "echo",
			wantArgs: []string{"hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, args := parseInput(tt.input)
			if cmd != tt.wantCmd {
				t.Errorf("parseInput() cmd = %v, want %v", cmd, tt.wantCmd)
			}
			if len(args) != len(tt.wantArgs) {
				t.Errorf("parseInput() args length = %v, want %v", len(args), len(tt.wantArgs))
				return
			}
			for i := range args {
				if args[i] != tt.wantArgs[i] {
					t.Errorf("parseInput() args[%d] = %v, want %v", i, args[i], tt.wantArgs[i])
				}
			}
		})
	}
}
