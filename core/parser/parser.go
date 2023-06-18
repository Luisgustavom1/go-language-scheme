package parser

import (
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

	return ast, index
}

func (a Ast) pretty() string {
	p := "("
	for _, value := range a {
		p += value.pretty()
		p += " "
	}

	return p + ")"
}


func (v AstValue) pretty() string {
	if v.Kind == LiteralValue {
		return v.Literal.Value
	}

	return v.List.pretty()
}