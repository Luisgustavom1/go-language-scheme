package lexer_test

import (
	"testing"

	"github.com/Luisgustavom1/go-language-scheme/core/lexer"
	"github.com/Luisgustavom1/go-language-scheme/core/models"
	"github.com/stretchr/testify/assert"
)

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
		cursor, token := lexer.LexIntegerToken([]rune(test.source), test.cursor)
		assert.Equal(t, cursor, test.expectedCursor)
		assert.Equal(t, token.Value, test.expectedValue)
		assert.Equal(t, token.Kind, models.IntegerToken)
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
		cursor, token := lexer.LexIdentifierToken([]rune(test.source), test.cursor)
		assert.Equal(t, cursor, test.expectedCursor)
		assert.Equal(t, token.Value, test.expectedValue)
		assert.Equal(t, token.Kind, models.IdentifierToken)
	}
}

func Test_lexer(t *testing.T) {
	tests := []struct {
		source string
		tokens []lexer.Token
	}{
		{
			source: " ( + 13 2  ) ",
			tokens: []lexer.Token{
				{
					Value:    "(",
					Kind:     models.SyntaxToken,
					Location: 1,
				},
				{
					Value:    "+",
					Kind:     models.IdentifierToken,
					Location: 3,
				},
				{
					Value:    "13",
					Kind:     models.IntegerToken,
					Location: 5,
				},
				{
					Value:    "2",
					Kind:     models.IntegerToken,
					Location: 8,
				},
				{
					Value:    ")",
					Kind:     models.SyntaxToken,
					Location: 11,
				},
			},
		},
	}

	for _, test := range tests {
		tokens := lexer.Lexer([]rune(test.source))

		for i, token := range tokens {
			assert.Equal(t, token.Value, tokens[i].Value)
			assert.Equal(t, token.Kind, tokens[i].Kind)
			assert.Equal(t, token.Location, tokens[i].Location)
		}
	}
}
