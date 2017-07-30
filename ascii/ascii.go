package main

import (
	"fmt"
	"github.com/knewbie/go-kit/terminal"
)

const (
	Gray = uint8(iota + 90)
	Red
	Green
	Yellow
	Blue
	Magenta
	//NRed      = uint8(31) // Normal
	EndColor = "\033[0m"
)

const (
	intro string = `
ASCII tool
You can see all visible ascii code
copyright @ kevin
`
	head  string = " \033[33mdec \033[32m| \033[31mhex \033[32m| \033[35mchar "
	head2 string = " dec | hex | char "
)

func main() {
	fmt.Println(intro)

	w := terminal.GetWidth()
	var charslen int = len(head2)
	col := int(w) / (charslen + 2)

	glen := col*charslen + (col-1)*2

	for q := 0; q < glen; q++ {
		fmt.Printf("\033[%dm-", Green)
	}

	fmt.Print("\n\033[37m")

	for t := 0; t < col; t++ {
		fmt.Print(head)
		if t < col-1 {
			fmt.Print("\033[34m||")
		} else {
			fmt.Println()
		}
	}

	for t := 0; t < col; t++ {
		for u := 0; u < charslen; u++ {
			fmt.Printf("\033[%dm-", 32)
		}
		if t < col-1 {
			fmt.Printf("\033[%dm||", 34)
		} else {
			fmt.Println()
		}
	}

	c := 127 - 32
	c = c / col //差值
	max := 32 + c
	for i := 32; i < max; i++ {
		for b := 0; b < col; b++ {
			z := i + b*c
			if z < 100 {
				fmt.Printf(" \033[33m%d  \033[32m|  \033[31m%x  \033[32m|  \033[35m%c  ", z, z, z)
			} else {
				fmt.Printf(" \033[33m%d \033[32m|  \033[31m%x  \033[32m|  \033[35m%c  ", z, z, z)
			}

			if b < col-1 {
				fmt.Printf("\033[34m||")
			}
		}
		fmt.Println()
	}

	for q := 0; q < glen; q++ {
		fmt.Printf("\033[32m-")
	}

	fmt.Print("\n\033[37m")

}
