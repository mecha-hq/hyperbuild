package model

import (
	"fmt"
	"os/exec"
)

var ErrBashCommandFailed = fmt.Errorf("command failed")

type Bash struct {
	Commands []string
}

func (d *Bash) Run() ([]string, error) {
	var outputs []string

	for _, command := range d.Commands {
		out, err := exec.Command("/bin/bash", "-c", command).Output()
		if err != nil {
			return outputs, fmt.Errorf("%w: %q - %w", ErrBashCommandFailed, command, err)
		}

		outputs = append(outputs, string(out))
	}

	return outputs, nil
}
