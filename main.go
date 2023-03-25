package main

import (
	"fmt"
	"os"

	"github.com/omissis/hyperbuild/internal/config"
	"github.com/omissis/hyperbuild/internal/model"
)

func main() {
	file := os.Args[1]
	m, err := config.ParseYAMLFile(file)
	if err != nil {
		panic(err)
	}

	res, err := model.Run(m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("done: %v", res)
}
