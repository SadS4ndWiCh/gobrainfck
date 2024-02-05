package main

import (
	"fmt"
	"gobrainfck/eval"
	"gobrainfck/lexer"
	"gobrainfck/parser"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: gobrainfck <file.bf>")
		os.Exit(1)
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read file '%s'", os.Args[1])
		os.Exit(1)
	}

	l := lexer.NewLexer(content)
	p := parser.NewParser(l)

	eval.Eval(p)
}
