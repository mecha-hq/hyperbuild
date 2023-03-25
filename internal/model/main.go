package model

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/natessilva/dag"
)

type Step struct {
	Name      string
	DependsOn []string
	Commands  []string
}

type Runner struct {
}

var (
	ErrCommandFailed = fmt.Errorf("command failed")
)

func Run(steps []Step) error {
	var r dag.Runner
	var outputs []string
	var mut sync.Mutex

	for _, step := range steps {
		r.AddVertex(step.Name, func() error {
			for _, command := range step.Commands {
				out, err := exec.Command(command).Output()
				if err != nil {
					return fmt.Errorf("%w: %q - %w", ErrCommandFailed, command, err)
				}

				mut.Lock()
				outputs = append(outputs, string(out))
				mut.Unlock()
			}

			return nil
		})
	}
}
