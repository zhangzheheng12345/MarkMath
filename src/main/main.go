package main

import (
    "fmt"
    "strings"
    "os"
    "elements"
)

func main() {
    template, err := os.ReadFile(".\template.html")
    if err != nil {
      fmt.Println(err)
      return
    }
    if len(os.Arg) > 1 {
      wholeFile, err := os.ReadFile(os.Args[1])
      if err != nil {
            fmt.Println(err)
      } else {
            Rend(ParseWhole(string(template),strings.Split(string(wholeFile),"\n")))
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
         Rend(ParseWhole(string(template),whole))
         fmt.Scanln(&input) /* wait enter */
      }
    }
}