package main

import (
	"github.com/k0kubun/pp"
	"log"
)

func ExampleSimpleParsing() {
	input := `
age = 21
name = "Osku rokkaa"

[address]
city = "Beverly Hills"
postal_code = 90210
`
	ini := INI{}
	err := Build(input, &ini)
	if err != nil {
		log.Fatal(err)
	}

	pp.Print(ini)
	// Output:
	// TODO
}
