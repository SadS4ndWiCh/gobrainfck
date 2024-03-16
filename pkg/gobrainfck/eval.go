package gobrainfck

import (
	"fmt"
	"os"
)

const MAX_MEMORY = 60

func Eval(p Parser) {
	memory := make([]int, MAX_MEMORY)
	ptr := 0

	program := p.Parse()

	for i := 0; i < len(program.Tokens); i++ {
		tk := program.Tokens[i]

		switch tk.Kind {
		case ADD:
			memory[ptr] += int(tk.Value)
		case SUB:
			memory[ptr] -= int(tk.Value)
		case FORWARD:
			if ptr >= MAX_MEMORY {
				fmt.Println("Stack overflow")
				os.Exit(1)
			}
			ptr += int(tk.Value)
		case BACKWARD:
			if ptr < 0 {
				fmt.Println("Stack underflow")
				os.Exit(1)
			}
			ptr -= int(tk.Value)
		case JUMP_IF:
			if memory[ptr] == 0 {
				i = int(tk.Value)
			}
		case JUMP_BACK_IF:
			if memory[ptr] != 0 {
				i = int(tk.Value)
			}
		case PRINT:
			for c := tk.Value; c > 0; c-- {
				fmt.Printf("%c", memory[ptr])
			}
		case INPUT:
			fmt.Scanf("%c", &memory[ptr])
		}
	}
}
