package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeLexingContext(source []rune) LexingContext {
	return LexingContext{
		Source:         source,
		SourceFileName: "<memory>",
	}
}

func Test_lexIntegerToken(t *testing.T) {
	tests := []struct {
		source         string
		cursor         int
		expectedValue  string
		expectedCursor int
	}{
		{
			"foo 123",
			4,
			"123",
			7,
		},
		{
			" 12",
			1,
			"12",
			3,
		},
		{
			"foo 12a 3",
			4,
			"12",
			6,
		},
	}

	for _, test := range tests {
		lc := makeLexingContext([]rune(test.source))
		cursor, token := lc.LexIntegerToken(test.cursor)

		assert.Equal(t, cursor, test.expectedCursor)
		assert.Equal(t, token.Value, test.expectedValue)
		assert.Equal(t, token.Kind, IntegerToken)
	}
}

func Test_lexIdentifierToken(t *testing.T) {
	tests := []struct {
		source         string
		cursor         int
		expectedValue  string
		expectedCursor int
	}{
		{
			"778 ab23 ",
			4,
			"ab23",
			8,
		},
		{
			"2 ab12 + ",
			2,
			"ab12",
			6,
		},
		{
			"function 3",
			0,
			"function",
			8,
		},
	}

	for _, test := range tests {
		lc := makeLexingContext([]rune(test.source))
		cursor, token := lc.LexIdentifierToken(test.cursor)

		assert.Equal(t, cursor, test.expectedCursor)
		assert.Equal(t, token.Value, test.expectedValue)
		assert.Equal(t, token.Kind, IdentifierToken)
	}
}

func Test_lexer(t *testing.T) {
	tests := []struct {
		source string
		tokens []Token
	}{
		{
			source: " ( + 13 2  ) ",
			tokens: []Token{
				{
					Value:    "(",
					Kind:     SyntaxToken,
					Location: 1,
				},
				{
					Value:    "+",
					Kind:     IdentifierToken,
					Location: 3,
				},
				{
					Value:    "13",
					Kind:     IntegerToken,
					Location: 5,
				},
				{
					Value:    "2",
					Kind:     IntegerToken,
					Location: 8,
				},
				{
					Value:    ")",
					Kind:     SyntaxToken,
					Location: 11,
				},
			},
		},
	}

	for _, test := range tests {
		lc := makeLexingContext([]rune(test.source))
		tokens := lc.Lexer()

		for i, token := range tokens {
			token.LexingContext = test.tokens[i].LexingContext
			assert.Equal(t, token, test.tokens[i])
		}
	}
}
