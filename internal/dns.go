package internal

import (
	"log"
	"os/exec"
	"strings"
)

func DnsGet(params ...string) (string, error) {
	_, err := exec.LookPath("dig")
	if err != nil {
		return "", err
	}

	cmd := exec.Command("dig", params...)
	stdout, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}
