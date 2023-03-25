package model

import (
	"fmt"
	"sync"

	"github.com/natessilva/dag"
)

type Manifest struct {
	Steps []Step
}

type Step struct {
	Name      string
	DependsOn []string `yaml:"depends_on"`

	Bash   *Bash
	Docker *Docker
}

func (s Step) Run() ([]string, error) {
	if s.Bash != nil {
		return s.Bash.Run()
	}

	if s.Docker != nil {
		return s.Docker.Run()
	}

	return nil, fmt.Errorf("no plugin found")
}

var (
	ErrStepFailed     = fmt.Errorf("step failed")
	ErrGraphRunFailed = fmt.Errorf("graph run failed")
)

func Run(m Manifest) ([]string, error) {
	var r dag.Runner
	var outputs []string
	var mut sync.Mutex

	for _, step := range m.Steps {
		step := step

		r.AddVertex(step.Name, func() error {
			out, err := step.Run()

			mut.Lock()
			outputs = append(outputs, out...)
			mut.Unlock()

			if err != nil {
				return fmt.Errorf("%w: %q - %w", ErrStepFailed, step.Name, err)
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
