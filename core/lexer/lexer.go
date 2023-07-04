package lexer

import (
	"os"
	"unicode"
)

func NewLexingContext(file string) LexingContext {
	sourceCode, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return LexingContext{
		SourceFileName: file,
		Source:         []rune(string(sourceCode)),
	}
}

func (lc LexingContext) Lexer() []Token {
	tokens := []Token{}
	var tokenTemp *Token

	cursor := 0
	for cursor < len(lc.Source) {
		cursor = skipWhitespace(lc.Source, cursor)
		if cursor == len(lc.Source) {
			break
		}

		cursor, tokenTemp = lc.LexSyntaxToken(cursor)
		if tokenTemp != nil {
			tokens = append(tokens, *tokenTemp)
			continue
		}

		cursor, tokenTemp = lc.LexIntegerToken(cursor)
		if tokenTemp != nil {
			tokens = append(tokens, *tokenTemp)
			continue
		}

		cursor, tokenTemp = lc.LexIdentifierToken(cursor)
		if tokenTemp != nil {
			tokens = append(tokens, *tokenTemp)
			continue
		}

		panic("Could nod lex")
	}

	return tokens
}

func skipWhitespace(sourceCode []rune, cursor int) int {
	for cursor < len(sourceCode) {
		if unicode.IsSpace(sourceCode[cursor]) {
			cursor++
			continue
		}

		break
	}

	return cursor
}

func (lc LexingContext) LexSyntaxToken(cursor int) (int, *Token) {
	if lc.Source[cursor] == '(' || lc.Source[cursor] == ')' {
		return cursor + 1, &Token{
			Value:         string([]rune{lc.Source[cursor]}),
			Kind:          SyntaxToken,
			Location:      cursor,
			LexingContext: lc,
		}
	}

	return cursor, nil
}

func (lc LexingContext) LexIntegerToken(cursor int) (int, *Token) {
	originalCursor := cursor

	values := []rune{}
	for cursor < len(lc.Source) {
		r := lc.Source[cursor]

		if r >= '0' && r <= '9' {
			values = append(values, r)
			cursor++
			continue
		}

		break
	}

	if len(values) == 0 {
		return originalCursor, nil
	}

	return cursor, &Token{
		Value:         string(values),
		Kind:          IntegerToken,
		Location:      originalCursor,
		LexingContext: lc,
	}
}

func (lc LexingContext) LexIdentifierToken(cursor int) (int, *Token) {
	originalCursor := cursor

	values := []rune{}
	for cursor < len(lc.Source) {
		r := lc.Source[cursor]

		if !unicode.IsSpace(r) && r != ')' {
			values = append(values, r)
			cursor++
			continue
		}

		break
	}

	if len(values) == 0 {
		return originalCursor, nil
	}

	return cursor, &Token{
		Value:         string(values),
		Kind:          IdentifierToken,
		Location:      originalCursor,
		LexingContext: lc,
	}
}
