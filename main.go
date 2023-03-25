package main

import (
	"errors"
	"fmt"

	"github.com/natessilva/dag"
)

func main() {
	var r dag.Runner

	r.AddVertex("one", func() error {
		fmt.Println("one and two will run in parallel before three")
		return nil
	})
	r.AddVertex("two", func() error {
		fmt.Println("one and two will run in parallel before three")
		return nil
	})
	r.AddVertex("three", func() error {
		fmt.Println("three will run before four")
		return errors.New("three is broken")
	})
	r.AddVertex("four", func() error {
		fmt.Println("four will never run")
		return nil
	})

	r.AddEdge("one", "three")
	r.AddEdge("two", "three")

	r.AddEdge("three", "four")

	fmt.Printf("the runner terminated with: %v\n", r.Run())
}
