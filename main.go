package main

import (
	"fmt"
	"org2fd/flowdock"
	"org2fd/orgmode"
	"os"
)

func main() {
	input := os.Stdin
	orgMode := &orgmode.OrgMode{}

	err := orgmode.ParseFile(input, orgMode)
	if err != nil {
		panic(err)
	}

	fmt.Print(flowdock.Format(orgMode))
}
