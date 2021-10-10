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
      wholeFile, err := string(os.ReadFile(os.Arg[1]))
      if err != nil {
            fmt.Println(err)
      } else {
            /* for windows, \r\n */
            ParseWhole(template,strings.Split(wholeFile,"\r\n")).rend()
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
         ParseWhole(template,li).rend()
         fmt.Scanln(&input) /* wait enter */
      }
    }
}