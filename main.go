package main

import (
	"fmt"

	"github.com/myugen/kata-videorecorder-go/kata"
)

func main() {
	dependency := kata.NewDependencyImpl()
	k := kata.New(dependency)

	fmt.Printf("the value is %d", k.Value())
}
