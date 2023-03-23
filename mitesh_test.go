package goshell

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestRunCommand0(t *testing.T) {
	expected_val := "hello"
	ret_val, _ := RunCommand("echo -n hello")
	if ret_val != expected_val {
		t.Errorf("RunCommand() = %q, want = %q", ret_val, expected_val)
	}
}

func TestRunCommand1(t *testing.T) {
	expected_val := "hello"
	ret_val, _ := RunCommand("echo hello")
	if ret_val != expected_val {
		t.Errorf("RunCommand() = %q, want = %q", ret_val, expected_val)
	}
}

func TestRunCommandPwd(t *testing.T) {
	expected_val, _ := os.Getwd()
	// expected_val += "\n"
	ret_val, _ := RunCommand("pwd")
	if ret_val != expected_val {
		t.Errorf("RunCommand() = %q, want = %q", ret_val, expected_val)
	}
}

func runCommand(command string) string {
	cmd := exec.Command("sh", "-c", command)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return ""
	}

	output := stdout.String()
	osplit := strings.Split(output, "\n")
	los := len(osplit)
	osjoin := strings.Join(osplit[0:los-1], "\n")
	return osjoin
}

func TestRunCommandPipe(t *testing.T) {
	eo := runCommand("ls -l")
	lines := SplitLines(eo)
	expected_val := len(lines)
	output, _ := RunCommand("ls -l | wc -l")
	ret_val, err := strconv.Atoi(output)
	if err != nil {
		fmt.Println(err)
	}
	if ret_val != expected_val {
		t.Errorf("RunCommandPipe() = %d, want = %d", ret_val, expected_val)
	}
}
