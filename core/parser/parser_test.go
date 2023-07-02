package parser

import (
	"testing"

	"github.com/Luisgustavom1/go-language-scheme/core/lexer"
	"github.com/stretchr/testify/assert"
)

func Test_parse(t *testing.T) {
	tests := []struct {
		input  string
		Pretty string
		output Ast
	}{
		{
			"(+ 1 2)",
			"(+ 1 2 )",
			Ast{
				{
					Kind: LiteralValue,
					Literal: &lexer.Token{
						Value: "+",
						Kind: lexer.SyntaxToken,
						Location: 0,
					},
				},
				{
					Kind:   LiteralValue,
					Literal: &lexer.Token{
						Value: "1",
						Kind: lexer.SyntaxToken,
						Location: 3,
					},
				},
				{
					Kind:   LiteralValue,
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
			"(+ 1 (- 12 9 ) )",
			Ast{
				{
					Kind: LiteralValue,
					Literal: &lexer.Token{Value: "+"},
				},
				{
					Kind: LiteralValue,
					Literal: &lexer.Token{Value: "1"},
				},
				{
					Kind: ListValue,
					List: &Ast{
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "-"},
						},
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "12"},
						},
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "9"},
						},
					},
				},
			},
		},
		{
			"(+ 1 (- 12 9) 12)",
			"(+ 1 (- 12 9 ) 12 )",
			Ast{
				{
					Kind: LiteralValue,
					Literal: &lexer.Token{Value: "+"},
				},
				{
					Kind: LiteralValue,
					Literal: &lexer.Token{Value: "1"},
				},
				{
					Kind: ListValue,
					List: &Ast{
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "-"},
						},
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "12"},
						},
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "9"},
						},
					},
				},
				{
					Kind: LiteralValue,
					Literal: &lexer.Token{Value: "12"},
				},
			},
		},
		{
			"(+ 1 (- 12 9) 4 (+ 76 1 ) 13)",
			"(+ 1 (- 12 9 ) 4 (+ 76 1 ) 13 )",
			Ast{
				{
					Kind: LiteralValue,
					Literal: &lexer.Token{Value: "+"},
				},
				{
					Kind: LiteralValue,
					Literal: &lexer.Token{Value: "1"},
				},
				{
					Kind: ListValue,
					List: &Ast{
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "-"},
						},
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "12"},
						},
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "9"},
						},
					},
				},
				{
					Kind: LiteralValue,
					Literal: &lexer.Token{Value: "4"},
				},
				{
					Kind: ListValue,
					List: &Ast{
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "+"},
						},
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "76"},
						},
						AstValue{
							Kind:    LiteralValue,
							Literal: &lexer.Token{Value: "1"},
						},
					},
				},
				{
					Kind: LiteralValue,
					Literal: &lexer.Token{Value: "13"},
				},
			},
		},
	}

	for _, test := range tests {
		tokens := lexer.Lexer([]rune(test.input))
		ast, _ := Parse(tokens, 0)

		assert.True(t, compareAst(ast, test.output))
		assert.Equal(t, ast.Pretty(), test.Pretty)
	}
}

func compareAst(a Ast, b Ast) bool {
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

func compareValue(a AstValue, b AstValue) bool {
	if a.Kind != b.Kind {
		return false
	}

	if a.Kind == LiteralValue {
		return a.Literal.Value == b.Literal.Value
	}

	return compareAst(*a.List, *b.List)
}
