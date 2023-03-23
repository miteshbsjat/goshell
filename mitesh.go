package goshell

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func RunCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	output := stdout.String()
	osplit := strings.Split(output, "\n")
	los := len(osplit)
	if osplit[los-1] == "" {
		osjoin := strings.Join(osplit[0:los-1], "\n")
		return osjoin, nil
	} else {
		return output, nil
	}
}

func SplitLines(str string) []string {
	return strings.Split(str, "\n")
}

func main1() {
	output, err := RunCommand("ls -l")
	if err != nil {
		fmt.Printf("Error running command: %s\n", err.Error())
		return
	}

	fmt.Printf("Command output:\n%s\n", output)
	for _, line := range SplitLines(output) {
		fmt.Println(line)
	}
}
