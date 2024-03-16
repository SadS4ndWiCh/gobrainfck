package gobrainfck

import (
	"strings"
)

type Lexer struct {
	source []byte
	pos    uint
	ch     rune
}

func NewLexer(source []byte) Lexer {
	l := Lexer{source: source}
	l.nextChar()

	return l
}

func (l *Lexer) nextChar() {
	if l.pos >= uint(len(l.source)) {
		l.ch = 0
		l.pos += 1
		return
	}

	l.ch = rune(l.source[l.pos])
	l.pos += 1
}

func (l *Lexer) skipUnknow() {
	for l.ch != 0 && !strings.Contains("+-<>[],.", string(l.ch)) {
		l.nextChar()
	}
}

func (l *Lexer) eat(char rune) {
	for l.ch == char {
		l.nextChar()
	}
}

func createToken(kind TokenKind, value uint) Token {
	return Token{Kind: kind, Value: value}
}

func (l *Lexer) NextToken() (tk Token) {
	l.skipUnknow()

	pos := l.pos

	switch l.ch {
	case '+':
		l.eat(l.ch)
		tk = createToken(ADD, l.pos-pos)
		return
	case '-':
		l.eat(l.ch)
		tk = createToken(SUB, l.pos-pos)
		return
	case '<':
		l.eat(l.ch)
		tk = createToken(BACKWARD, l.pos-pos)
		return
	case '>':
		l.eat(l.ch)
		tk = createToken(FORWARD, l.pos-pos)
		return
	case '[':
		tk = createToken(JUMP_IF, 0)
	case ']':
		tk = createToken(JUMP_BACK_IF, 0)
	case ',':
		l.eat(l.ch)
		tk = createToken(INPUT, l.pos-pos)
		return
	case '.':
		l.eat(l.ch)
		tk = createToken(PRINT, l.pos-pos)
		return
	case 0:
		tk = createToken(EOF, 0)
	default:
		tk = createToken(INVALID, 0)
	}

	l.nextChar()
	return
}
