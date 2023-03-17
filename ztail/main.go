package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	pck := os.Args[1:]
	if len(pck) == 0 {
		return
	}
	res := 0
	for i := 0; i < len(pck); i++ {
		if pck[1] == "%" {
			res = piscine.Atoi(pck[0]) % piscine.Atoi(pck[2])

			for _, c := range string(res) {
				z01.PrintRune(c)
			}
		}
	}
}
