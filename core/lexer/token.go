package lexer

type Token struct {
	Value    string    `json:"value"`
	Kind     TokenKind `json:"kind"`
	Location int       `json:"location"`
}

type TokenKind uint

const (
	// "(", ")"
	SyntaxToken TokenKind = iota
	// "1", "34"
	IntegerToken
	// "+", "define"
	IdentifierToken
)
