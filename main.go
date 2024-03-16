package main

import (
	"fmt"
	"gobrainfck/cmd"
	"os"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
