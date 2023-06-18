package parser

import (
	"github.com/Luisgustavom1/go-language-scheme/core/lexer"
)

type Ast []AstValue

type AstValue struct {
	Kind    valueKind
	Literal *lexer.Token
	List    *Ast
}

type valueKind uint

const (
	LiteralValue valueKind = iota
	ListValue
)