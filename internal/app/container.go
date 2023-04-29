package app

import (
	"fmt"
)

var ErrCannotCreateContainer = fmt.Errorf("cannot create container")

type ContainerFactoryFunc func() (*Container, error)

func NewDefaultParameters() Parameters {
	return Parameters{}
}

type Parameters struct {
	Versions Versions
}

type services struct{}

func NewContainer(versions Versions) *Container {
	ctr := Container{
		Parameters: NewDefaultParameters(),
	}

	ctr.Versions = versions

	return &ctr
}

type Container struct {
	Parameters
	services
}
