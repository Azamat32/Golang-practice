package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	pck_name := os.Args[0]
	exec_name := ""

	for i := len(pck_name) - 1; i >= 0; i-- {
		if pck_name[i] == '/' {
			exec_name = pck_name[i+1:]
			break
		}
	}
	for _, c := range exec_name {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
