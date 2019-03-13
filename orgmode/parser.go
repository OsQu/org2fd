package orgmode

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"io"
)

var orgModeLexer = lexer.Must(lexer.Regexp(
	`(\s+)` +
		`|(?P<Dot>\*)` +
		`|(?P<String>.*)`,
))

type OrgMode struct {
	Headlines []*Headline `@@*`
}

type Headline struct {
	Line *Line `@@`
}

type Line struct {
	String *string `"*" @String`
}

// Parse parses input to orgMode representing AST
func Parse(input string, orgMode *OrgMode) error {
	parser := buildParser()
	return parser.ParseString(input, orgMode)
}

// ParseFile parses input from file to orgMode representing AST
func ParseFile(input io.Reader, orgMode *OrgMode) error {
	parser := buildParser()
	return parser.Parse(input, orgMode)
}

func buildParser() *participle.Parser {
	parser, err := participle.Build(&OrgMode{},
		participle.Lexer(orgModeLexer),
	)

	if err != nil {
		panic(err)
	}

	return parser
}
