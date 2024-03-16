package gobrainfck

const (
	EOF     = "EOF"
	INVALID = "INVALID"

	ADD      = "ADD"
	SUB      = "SUB"
	FORWARD  = "FORWARD"
	BACKWARD = "BACKWARD"

	JUMP_IF      = "JUMP_IF"
	JUMP_BACK_IF = "JUMP_BACK_IF"

	PRINT = "PRINT"
	INPUT = "INPUT"
)

type TokenKind string
type Token struct {
	Kind  TokenKind
	Value uint
}
