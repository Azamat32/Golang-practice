package main

import (
	"fmt"
	"os"
)

func main() {
	pck_name := os.Args[1:]

	for i := 0; i < len(pck_name); i++ {
		if pck_name[i] == "01" || pck_name[i] == "galaxy" || pck_name[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
