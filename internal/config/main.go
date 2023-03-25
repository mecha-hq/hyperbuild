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

func ParseYAMLFile(path string) (model.Manifest, error) {
	content, err := os.ReadFile(path)
	var manifest model.Manifest
	if err != nil {
		return manifest, fmt.Errorf("%w: %w", ErrCannotReadConfigFile, err)
	}

	if err := yaml.Unmarshal([]byte(content), &manifest); err != nil {
		return manifest, fmt.Errorf("%w: %w", ErrCannotParseConfigFile, err)
	}

	return manifest, nil
}
