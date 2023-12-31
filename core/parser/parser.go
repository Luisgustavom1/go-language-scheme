package parser

import (
	"os"

	"github.com/Luisgustavom1/go-language-scheme/core/lexer"
)

func Parse(tokens []lexer.Token, index int) (Ast, int) {
	var ast Ast

	index++
	for index < len(tokens) {
		token := tokens[index]

		if token.Kind == lexer.SyntaxToken && token.Value == "(" {
			child, nextIndex := Parse(tokens, index)

			ast = append(ast, AstValue{
				Kind:    ListValue,
				Literal: &token,
				List:    &child,
			})
			index = nextIndex
			continue
		}

		if token.Kind == lexer.SyntaxToken && token.Value == ")" {
			return ast, index + 1
		}

		ast = append(ast, AstValue{
			Kind:    LiteralValue,
			Literal: &token,
		})
		index++
	}

	lastToken := tokens[index-1]
	if lastToken.Kind == lexer.SyntaxToken && lastToken.Value != ")" {
		lastToken.Debug("Expected closing paren")
		os.Exit(1)
	}

	return ast, index
}

func (a Ast) Pretty() string {
	p := "("
	for _, value := range a {
		p += value.Pretty()
		p += " "
	}

	return p + ")"
}

func (v AstValue) Pretty() string {
	if v.Kind == LiteralValue {
		return v.Literal.Value
	}

	return v.List.Pretty()
}
