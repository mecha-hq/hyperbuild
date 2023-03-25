package model

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/natessilva/dag"
)

type Manifest struct {
	Steps []Step
}

type Step struct {
	Name      string
	DependsOn []string `yaml:"depends_on"`
	Commands  []string
}

type Runner struct {
}

var (
	ErrCommandFailed  = fmt.Errorf("command failed")
	ErrGraphRunFailed = fmt.Errorf("graph run failed")
)

func Run(steps []Step) ([]string, error) {
	var r dag.Runner
	var outputs []string
	var mut sync.Mutex

	for _, step := range steps {
		step := step
		r.AddVertex(step.Name, func() error {
			for _, command := range step.Commands {
				fmt.Printf("running command: %q\n", command)
				command := command
				// args := strings.Split(command, " ")
				out, err := exec.Command("/bin/bash", "-c", command).Output()
				if err != nil {
					return fmt.Errorf("%w: %q - %w", ErrCommandFailed, command, err)
				}

				mut.Lock()
				outputs = append(outputs, string(out))
				mut.Unlock()
			}

			return nil
		})

		for _, don := range step.DependsOn {
			r.AddEdge(don, step.Name)
		}
	}

	if err := r.Run(); err != nil {
		return outputs, fmt.Errorf("%w: %w", ErrGraphRunFailed, err)
	}

	return outputs, nil
}
