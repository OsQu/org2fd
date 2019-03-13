package orgmode

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func lex(input string) []lexer.Token {
	parser, err := participle.Build(&OrgMode{},
		participle.Lexer(orgModeLexer),
	)
	if err != nil {
		panic(err)
	}

	lexer, _ := parser.Lex(strings.NewReader(input))

	return lexer
}

func TestLexerSingleLine(t *testing.T) {
	input := `* Headline`
	tokens := lex(input)

	assert.Equal(t, 3, len(tokens))
}

func TestLexerMultipleLines(t *testing.T) {
	input := `
* Headline
* Another headline
`
	tokens := lex(input)

	assert.Equal(t, 5, len(tokens))
}
