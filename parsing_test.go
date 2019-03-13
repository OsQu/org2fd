package main

import (
	"fmt"
	"org2fd/flowdock"
	"org2fd/orgmode"
)

func parseInput(input string) string {
	orgMode := &orgmode.OrgMode{}
	err := orgmode.Parse(input, orgMode)

	if err != nil {
		panic(err)
	}

	return flowdock.Format(orgMode)
}

func ExampleSimpleParsing() {
	input := `* This is a headline`
	fmt.Println(parseInput(input))

	// Output:
	// - This is a headline
}

func ExampleParseTwoLines() {
	input := `
* This is a headline
* This is another headline
`
	fmt.Println(parseInput(input))

	// Output:
	// - This is a headline
	// - This is another headline
}
