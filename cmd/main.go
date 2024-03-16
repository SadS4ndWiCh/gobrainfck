package cmd

import (
	"errors"
	"fmt"
	"gobrainfck/pkg/gobrainfck"
	"os"
)

func Run() error {
	if len(os.Args) == 1 {
		return errors.New("usage: gobrainfck <file.bf>")
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		return fmt.Errorf("failed to read file '%s'", os.Args[1])
	}

	l := gobrainfck.NewLexer(content)
	p := gobrainfck.NewParser(l)

	gobrainfck.Eval(p)
	return nil
}
