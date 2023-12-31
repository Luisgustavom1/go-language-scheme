package main

import (
	"fmt"
	"os"

	"github.com/Luisgustavom1/go-language-scheme/core/interpreter"
	"github.com/Luisgustavom1/go-language-scheme/core/lexer"
	"github.com/Luisgustavom1/go-language-scheme/core/parser"
)

func main() {
	lc := lexer.NewLexingContext(os.Args[1])

	tokens := lc.Lexer()

	var parseIndex int
	var ast = parser.Ast{
		{
			Kind:    parser.LiteralValue,
			Literal: &lexer.Token{Value: "begin"},
		},
	}

	for parseIndex < len(tokens) {
		childAst, nextIndex := parser.Parse(tokens, parseIndex)
		ast = append(ast, parser.AstValue{
			Kind: parser.ListValue,
			List: &childAst,
		})

		parseIndex = nextIndex
	}

	// Optimizations
	// Type checking
	// Are made in this step, over the AST

	interpreter.InitializeBuiltIns()
	value := interpreter.InitAstWalk(ast, map[string]any{})
	fmt.Println(value)
}
