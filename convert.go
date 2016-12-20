package convert

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
)

type converCmd interface {
	Version() string
	Execute(args []string) (string, error)
}

const CONVERT = "convert"

func getConvertCmd() (converCmd, error) {
	commandPath, err := lookPath()
	if err != nil {
		return nil, err
	}
	v, err := version()
	if err != nil {
		return nil, err
	}
	if regexp.MustCompile(`^6.9`).Match([]byte(v)) {
		cmd := convert0609{version: v, commandPath: commandPath}
		return cmd, nil
	}
	err = fmt.Errorf("No convert version is found, supporting 6.9~")
	return nil, err
}
func version() (string, error) {
	v, err := execConvertCmdWithStderr("-version")
	if err != nil {
		return v, err
	}
	exp := regexp.MustCompile(`^Version: ImageMagick ([0-9\\.]+)`)
	matches := exp.FindStringSubmatch(v)
	if len(matches) < 2 {
		err = fmt.Errorf("tesseract version not found: response is `%s`", v)
		return "", err
	}
	v = matches[1]
	return v, nil
}

func execConvertCmdWithStderr(opt string) (string, error) {
	cmd := exec.Command(CONVERT, opt)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdout.String(), nil
}

func lookPath() (commandPath string, err error) {
	return exec.LookPath(CONVERT)

}
