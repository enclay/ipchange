package internal

import (
	"os/exec"
	"strings"
)

func ExecStr(cmdstr string) (string, error) {
	arr := strings.Fields(cmdstr)
	cmd := exec.Command(arr[0], arr[1:]...)

	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}
