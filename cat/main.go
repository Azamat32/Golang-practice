package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	pck_name := os.Args[1:]
	stringErr := "ERROR: "

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, _ := ioutil.ReadAll(os.Stdin)
		for _, b := range data {
			z01.PrintRune(rune(b))
		}
	}
	if 0 == len(pck_name) {
		return
	} else {
		for _, c := range os.Args[1:] {
			file, err := os.Open(c)
			if err != nil {
				for _, z := range stringErr {
					z01.PrintRune(rune(z))
				}
				for _, z := range err.Error() {
					z01.PrintRune(rune(z))
				}
				z01.PrintRune('\n')

				os.Exit(1)
			} else {
				data, err := ioutil.ReadAll(file)
				if err != nil {
					for _, z := range stringErr {
						z01.PrintRune(rune(z))
					}
					for _, x := range err.Error() {
						z01.PrintRune(rune(x))
					}
					z01.PrintRune('\n')

				} else {
					for _, b := range data {
						z01.PrintRune(rune(b))
					}
				}
			}

		}
	}
}
