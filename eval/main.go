package eval

import (
	"fmt"
	"gobrainfck/lexer"
	"gobrainfck/parser"
	"os"
)

const MAX_MEMORY = 60

func Eval(p parser.Parser) {
	memory := make([]int, MAX_MEMORY)
	ptr := 0

	program := p.Parse()

	for i := 0; i < len(program.Tokens); i++ {
		tk := program.Tokens[i]

		switch tk.Kind {
		case lexer.ADD:
			memory[ptr] += int(tk.Value)
		case lexer.SUB:
			memory[ptr] -= int(tk.Value)
		case lexer.FORWARD:
			if ptr >= MAX_MEMORY {
				fmt.Println("Stack overflow")
				os.Exit(1)
			}
			ptr += int(tk.Value)
		case lexer.BACKWARD:
			if ptr < 0 {
				fmt.Println("Stack underflow")
				os.Exit(1)
			}
			ptr -= int(tk.Value)
		case lexer.JUMP_IF:
			if memory[ptr] == 0 {
				i = int(tk.Value)
			}
		case lexer.JUMP_BACK_IF:
			if memory[ptr] != 0 {
				i = int(tk.Value)
			}
		case lexer.PRINT:
			for c := tk.Value; c > 0; c-- {
				fmt.Printf("%c", memory[ptr])
			}
		case lexer.INPUT:
			fmt.Scanf("%c", &memory[ptr])
		}
	}
}
