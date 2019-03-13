package orgmode

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
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
	parser, err := participle.Build(&OrgMode{},
		participle.Lexer(orgModeLexer),
	)

	if err != nil {
		panic(err)
	}

	return parser.ParseString(input, orgMode)
}
