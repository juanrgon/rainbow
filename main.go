package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/juanrgon/prism"
)

func main() {
	fmt.Println(rainbow(strings.Join(os.Args[1:], " ")))
}

func rainbow(s string) (r prism.DecoratedString) {
	ds := prism.DecoratedString(s)
	for i, c := range ds {
		switch c := string(c); i % 5 {
		case 0:
			r += prism.InRed(c)
		case 1:
			r += prism.InYellow(c)
		case 2:
			r += prism.InGreen(c)
		case 3:
			r += prism.InCyan(c)
		case 4:
			r += prism.InMagenta(c)
		}
	}
	return
}
