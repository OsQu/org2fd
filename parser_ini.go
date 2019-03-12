package main

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
)

var iniLexer = lexer.Must(lexer.Regexp(
	`(?m)` +
		`(\s+)` +
		`|(^[#;].*$)` +
		`|(?P<Ident>[a-zA-Z][a-zA-Z_\d]*)` +
		`|(?P<String>"(?:\\.|[^"])*")` +
		`|(?P<Float>\d+(?:\.\d+)?)` +
		`|(?P<Punct>[][=])`,
))

type INI struct {
	Properties []*Property `@@*`
	Sections   []*Section  `@@*`
}

type Section struct {
	Identifier string      `"[" @Ident "]"`
	Properties []*Property `@@*`
}

type Property struct {
	Key   string `@Ident "="`
	Value *Value `@@`
}

type Value struct {
	String *string  `  @String`
	Number *float64 `| @Float`
}

// Build parses the input
func Build(input string, ini *INI) error {
	parser, err := participle.Build(&INI{},
		participle.Lexer(iniLexer),
		participle.Unquote("String"),
	)
	if err != nil {
		panic(err)
	}

	return parser.ParseString(input, ini)
}
