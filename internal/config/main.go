package config

import (
	"fmt"
	"os"

	"github.com/omissis/hyperbuild/internal/model"
	"gopkg.in/yaml.v3"
)

var (
	ErrCannotReadConfigFile  = fmt.Errorf("cannot read config file")
	ErrCannotParseConfigFile = fmt.Errorf("cannot parse config file")
)

func ParseYAMLFile(path string) ([]model.Step, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotReadConfigFile, err)
	}

	if err := yaml.Unmarshal([]byte(content), &model.Step{}); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotParseConfigFile, err)
	}

	return nil, nil
}
