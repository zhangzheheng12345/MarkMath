package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zhangzheheng12345/MarkMath/parser"
)

func main() {
	if len(os.Args) > 1 {
		path, _ := filepath.Abs(os.Args[1])
		wholeFile, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		} else {
			parser.Rend(parser.ParseWhole(strings.Split(string(wholeFile), "\n")))
		}
	} else {
		var input string = ""
		var whole []string
		reader := bufio.NewReader(os.Stdin)
		for input != "\\quit\r" && input != "\\rend\r" {
			fmt.Print(len(whole)+1, "> ")
			input, _ = reader.ReadString('\n')
			input = input[:len(input)-1]
			whole = append(whole, input)
			if input == "\\quit" && input == "\\rend" {
				break
			}
		}
		if input == "\\rend" || /* For Windows */ input == "\\rend\r" {
			whole = whole[:len(whole)-1]
			parser.Rend(parser.ParseWhole(whole))
		}
	}
	a := ""
	fmt.Scanln(&a)
}
