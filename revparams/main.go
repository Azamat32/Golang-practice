package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	pck_name := os.Args
	lendth := 0
	for c := range pck_name {
		lendth = c
	}
	for i := lendth; i >= 1; i-- {
		for _, c := range pck_name[i] {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
