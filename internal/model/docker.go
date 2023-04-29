package model

import (
	"fmt"
	"os/exec"
)

type Docker struct {
	File      string
	Tags      []string
	BuildArgs []string `yaml:"build_args"`
}

func (d Docker) Run() ([]string, error) {
	args := []string{"build", "-f", d.File}
	for _, tag := range d.Tags {
		args = append(args, "--tag", tag)
	}

	for _, arg := range d.BuildArgs {
		args = append(args, "--build-arg", arg)
	}

	args = append(args, ".")
	fmt.Println(args)

	out, err := exec.Command("docker", args...).Output()
	if err != nil {
		return nil, err
	}

	return []string{string(out)}, nil
}
