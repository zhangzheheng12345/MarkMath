package main

import (
    "fmt"
    "strings"
    "os"
    "elements"
)

func main() {
    if len(os.Arg) > 1 {
      var wholeFile string, err = string(os.ReadFile(os.Arg[1]))
      if err != nil {
            fmt.Println(err)
      } else {
            /* for windows */
            ParseWhole(strings.Split(wholeFile,"\r\n")).rend()
      }
      fmt.Scanln(&wholeFile)
    } else {
      var input string
      var whole []string
      for input != "\quit" && input != "\rend"{
         fmt.Print("> ")
         fmt.Scanln(&input)
         whole = append(whole,input)
      }
      if input == "\rend" {
         elements.Writeln("<!DOCTYPE html>")
         elements.Writeln("<html>")
         elements.Writeln("<head>")
         elements.Writeln("<style>")
         elements.Writeln("div{")
         elements.Writeln("   display:flex;")
         elements.Writeln("   justify-content:flex-start;")
         elements.Writeln("   align-items:center;")
         elements.Writeln("   padding:3px;")
         elements.Writeln("}")
         elements.Writeln("</style>")
         elements.Writeln("</head>")
         elements.Writeln("<body>")
         ParseWhole(li).rend()
         elements.Writeln("</body>")
         elements.Writeln("</html>")
         fmt.Scanln(&input) /* wait enter */
      }
    }
}