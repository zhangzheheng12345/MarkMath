package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	template, err := os.ReadFile(".\template.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(os.Args) > 1 {
		wholeFile, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
		} else {
			/* for windows, \r\n */
			x := ParseWhole(string(template), strings.Split(string(wholeFile), "\n"))
			x.rend()
		}
		fmt.Scanln(&wholeFile)
	} else {
		var input string
		var whole []string
		for input != "\\quit" && input != "\rend" {
			fmt.Print("> ")
			fmt.Scanln(&input)
			whole = append(whole, input)
		}
		if input == "\rend" {
			ParseWhole(string(template), whole).rend()
			fmt.Scanln(&input) /* wait enter */
		}
	}
}
