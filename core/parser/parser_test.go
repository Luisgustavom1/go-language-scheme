package parser_test

import (
	"testing"

	"github.com/Luisgustavom1/go-language-scheme/core/lexer"
	"github.com/Luisgustavom1/go-language-scheme/core/parser"
	"github.com/stretchr/testify/assert"
)

func Test_parse(t *testing.T) {
	tests := []struct {
		input  string
		output parser.Ast
	}{
		{
			"(+ 1 2)",
			parser.Ast{
				{
					Kind: parser.LiteralValue,
					Literal: &lexer.Token{
						Value: "+",
						Kind: lexer.SyntaxToken,
						Location: 0,
					},
				},
				{
					Kind:   parser.LiteralValue,
					Literal: &lexer.Token{
						Value: "1",
						Kind: lexer.SyntaxToken,
						Location: 3,
					},
				},
				{
					Kind:   parser.LiteralValue,
					Literal: &lexer.Token{
						Value: "2",
						Kind: lexer.SyntaxToken,
						Location: 5,
					},
				},
			},
		},
		{
			"(+ 1 (- 12 9))",
			parser.Ast{
				{
					Kind: parser.LiteralValue,
					Literal: &lexer.Token{Value: "+"},
				},
				{
					Kind: parser.LiteralValue,
					Literal: &lexer.Token{Value: "1"},
				},
				{
					Kind: parser.ListValue,
					List: &parser.Ast{
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "-"},
						},
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "12"},
						},
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "9"},
						},
					},
				},
			},
		},
		{
			"(+ 1 (- 12 9) 12)",
			parser.Ast{
				{
					Kind: parser.LiteralValue,
					Literal: &lexer.Token{Value: "+"},
				},
				{
					Kind: parser.LiteralValue,
					Literal: &lexer.Token{Value: "1"},
				},
				{
					Kind: parser.ListValue,
					List: &parser.Ast{
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "-"},
						},
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "12"},
						},
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "9"},
						},
					},
				},
				{
					Kind: parser.LiteralValue,
					Literal: &lexer.Token{Value: "12"},
				},
			},
		},
		{
			"(+ 1 (- 12 9) 4 (+ 76 1 ) 13)",
			parser.Ast{
				{
					Kind: parser.LiteralValue,
					Literal: &lexer.Token{Value: "+"},
				},
				{
					Kind: parser.LiteralValue,
					Literal: &lexer.Token{Value: "1"},
				},
				{
					Kind: parser.ListValue,
					List: &parser.Ast{
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "-"},
						},
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "12"},
						},
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "9"},
						},
					},
				},
				{
					Kind: parser.LiteralValue,
					Literal: &lexer.Token{Value: "4"},
				},
				{
					Kind: parser.ListValue,
					List: &parser.Ast{
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "+"},
						},
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "76"},
						},
						parser.AstValue{
							Kind:    parser.LiteralValue,
							Literal: &lexer.Token{Value: "1"},
						},
					},
				},
				{
					Kind: parser.LiteralValue,
					Literal: &lexer.Token{Value: "13"},
				},
			},
		},
	}

	for _, test := range tests {
		tokens := lexer.Lexer([]rune(test.input))
		ast, _ := parser.Parse(tokens, 0)

		assert.True(t, compareAst(ast, test.output))
	}
}

func compareAst(a parser.Ast, b parser.Ast) bool {
	if len(a) != len(b) {
		return false
	}

	for i, aValue := range a {
		if !compareValue(aValue, b[i]) {
			return false
		}
	}

	return true
}

func compareValue(a parser.AstValue, b parser.AstValue) bool {
	if a.Kind != b.Kind {
		return false
	}

	if a.Kind == parser.LiteralValue {
		return a.Literal.Value == b.Literal.Value
	}

	return compareAst(*a.List, *b.List)
}
