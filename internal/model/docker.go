package model

type Docker struct {
	File      string
	Tags      []string
	BuildArgs []string `yaml:"build_args"`
}

func (d Docker) Run() ([]string, error) {
	panic("not implemented")
}
