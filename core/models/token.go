package models

type Token struct {
	Value    string    `json:"value"`
	Kind     TokenKind `json:"kind"`
	Location int       `json:"location"`
}

type TokenKind uint

const (
	SyntaxToken TokenKind = iota
	IntegerToken
	IdentifierToken
)
