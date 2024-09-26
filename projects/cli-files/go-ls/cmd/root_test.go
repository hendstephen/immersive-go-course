package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func setOutputs(cmd *cobra.Command) (*bytes.Buffer, *bytes.Buffer) {
	out := bytes.NewBufferString("")
	err := bytes.NewBufferString("")

	cmd.SetOut(out)
	cmd.SetErr(err)

	return out, err
}

func TestExecuteTooManyArgs(t *testing.T) {
	args := []string{"file1", "file2"}
	cmd := NewCmd()
	cmd.SetArgs(args)
	setOutputs(cmd)

	err := cmd.Execute()
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestExecuteDir(t *testing.T) {
	args := []string{"../assets"}

	cmd := NewCmd()
	cmd.SetArgs(args)
	out, _ := setOutputs(cmd)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	expected := "dew.txt\nfor_you.txt\nrain.txt\n"
	if out.String() != expected {
		t.Fatalf("Expected %v, got %v", expected, out.String())
	}
}

func TestExecuteFile(t *testing.T) {
	filename := "root.go"
	args := []string{filename}

	cmd := NewCmd()
	cmd.SetArgs(args)
	out, _ := setOutputs(cmd)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	expected := fmt.Sprintf("%v\n", filename)
	if out.String() != expected {
		t.Fatalf("Expected %v, got %v", expected, out.String())
	}
}

func TestExecuteNoArgs(t *testing.T) {
	cmd := NewCmd()
	out, _ := setOutputs(cmd)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	expected := "root.go\nroot_test.go\n"
	if out.String() != expected {
		t.Fatalf("Expected %v, got %v", expected, out.String())
	}
}
