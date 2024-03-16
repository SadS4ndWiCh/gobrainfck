package gobrainfck

import (
	"fmt"
	"os"
)

type Parser struct {
	lexer     Lexer
	currToken Token
	peekToken Token
}

type Program struct {
	Tokens []Token
}

func NewParser(lexer Lexer) Parser {
	p := Parser{lexer: lexer}
	p.nextToken()
	p.nextToken()

	return p
}

func (p Parser) currTokenIs(kind TokenKind) bool { return p.currToken.Kind == kind }

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) Parse() Program {
	program := Program{}
	jumps := make([]uint, 0)

	for i := uint(0); !p.currTokenIs(EOF); i++ {
		switch p.currToken.Kind {
		case JUMP_IF:
			jumps = append(jumps, i)
		case JUMP_BACK_IF:
			if len(jumps) == 0 {
				fmt.Println("Unopened jump")
				os.Exit(1)
			}

			addr := jumps[len(jumps)-1]
			jumps = jumps[:len(jumps)-1]

			program.Tokens[addr].Value = i
			p.currToken.Value = addr
		}

		program.Tokens = append(program.Tokens, p.currToken)
		p.nextToken()
	}

	if len(jumps) > 0 {
		fmt.Println("Some jump aren't closed")
		os.Exit(1)
	}

	return program
}
