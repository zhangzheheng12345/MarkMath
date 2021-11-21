package main

import (
	"elements"
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

func main() {
	template, err := os.ReadFile("template.html")
	if err != nil {
		fmt.Println(err)
	} else if len(os.Args) > 1 {
		path, _ := filepath.Abs(os.Args[1])
		wholeFile, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		} else {
			elements.Rend(ParseWhole(string(template), strings.Split(string(wholeFile), "\n")))
		}
	} else {
		var input string = ""
		var whole []string
		for input != "\\quit" && input != "\\rend" {
			fmt.Print("> ")
			fmt.Scanln(&input)
			whole = append(whole, input)
		}
		if input == "\\rend" {
			elements.Rend(ParseWhole(string(template), whole))
		}
	}
	a := ""
	fmt.Scanln(&a)
}
