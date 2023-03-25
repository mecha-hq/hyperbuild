package model

import (
	"fmt"
	"sync"

	"github.com/natessilva/dag"
)

type Manifest struct {
	Steps []Step
}

type Step interface {
	Name() string
	DependsOn() []string
	Run() ([]byte, error)
}

var (
	ErrCommandFailed  = fmt.Errorf("command failed")
	ErrGraphRunFailed = fmt.Errorf("graph run failed")
)

func Run(m Manifest) ([]string, error) {
	var r dag.Runner
	var outputs []string
	var mut sync.Mutex

	for _, step := range m.Steps {
		step := step

		r.AddVertex(step.Name(), func() error {

			out, err := step.Run()
			if err != nil {
				return fmt.Errorf("%w: %q - %w", ErrCommandFailed, step.Name(), err)
			}

			mut.Lock()
			outputs = append(outputs, string(out))
			mut.Unlock()

			r.Run()

			return nil
		})

		for _, don := range step.DependsOn() {
			r.AddEdge(don, step.Name())
		}
	}

	if err := r.Run(); err != nil {
		return outputs, fmt.Errorf("%w: %w", ErrGraphRunFailed, err)
	}

	return outputs, nil
}
