package lexer

import (
	"unicode"
)

func Lexer(sourceCode []rune) []Token {
	tokens := []Token{}
	var tokenTemp *Token

	cursor := 0
	for cursor < len(sourceCode) {
		cursor = skipWhitespace(sourceCode, cursor)
		if cursor == len(sourceCode) {
			break
		}

		cursor, tokenTemp = LexSyntaxToken(sourceCode, cursor)
		if tokenTemp != nil {
			tokens = append(tokens, *tokenTemp)
			continue
		}

		cursor, tokenTemp = LexIntegerToken(sourceCode, cursor)
		if tokenTemp != nil {
			tokens = append(tokens, *tokenTemp)
			continue
		}

		cursor, tokenTemp = LexIdentifierToken(sourceCode, cursor)
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

func LexSyntaxToken(sourceCode []rune, cursor int) (int, *Token) {
	if sourceCode[cursor] == '(' || sourceCode[cursor] == ')' {
		return cursor + 1, &Token{
			Value:    string([]rune{sourceCode[cursor]}),
			Kind:     SyntaxToken,
			Location: cursor,
		}
	}

	return cursor, nil
}

func LexIntegerToken(sourceCode []rune, cursor int) (int, *Token) {
	originalCursor := cursor

	values := []rune{}
	for cursor < len(sourceCode) {
		r := sourceCode[cursor]

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
		Value:    string(values),
		Kind:     IntegerToken,
		Location: originalCursor,
	}
}

func LexIdentifierToken(sourceCode []rune, cursor int) (int, *Token) {
	originalCursor := cursor

	values := []rune{}
	for cursor < len(sourceCode) {
		r := sourceCode[cursor]

		if !unicode.IsSpace(r) {
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
		Value:    string(values),
		Kind:     IdentifierToken,
		Location: originalCursor,
	}
}
