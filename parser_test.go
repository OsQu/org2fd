package main

import (
	"github.com/k0kubun/pp"
)

func ExampleSimpleParsing() {
	input := `* This is a headline`
	orgMode := OrgMode{}
	err := Parse(input, &orgMode)

	if err != nil {
		panic(err)
	}

	pp.Print(orgMode)

	// Output:
	// TODO
}
