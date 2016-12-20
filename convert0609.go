package convert

import (
	"bytes"
	"fmt"
	"os/exec"
)

type convert0609 struct {
	version        string
	resultFilePath string
	commandPath    string
}

func (c convert0609) Version() string {
	return c.version
}

func (c convert0609) Execute(params []string) (string, error) {
	var args []string
	args = append(args, "-density")
	args = append(args, "200")

	args = append(args, params[0])

	if len(params) > 1 {
		args = append(args, params[1])
	}
	cmd := exec.Command(CONVERT, args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf(stderr.String())
	}

	return "done", nil
}
