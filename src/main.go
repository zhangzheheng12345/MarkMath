package main

import (
    "fmt"
    "strings"
    "elements"
)

func main() {
    var input string
    fmt.Scanln(&input)
    for input != "quit" {
        Writeln("<!DOCTYPE html>")
        Writeln("<html>")
        Writeln("<head>")
        Writeln("<style>")
        Writeln("div{")
        Writeln("   display:flex;")
        Writeln("   justify-content:flex-start;")
        Writeln("   align-items:center;")
        Writeln("   padding:3px;")
        Writeln("}")
        Writeln("div#main{")
        Writeln("   padding:7px;")
        Writeln("}")
        Writeln("</style>")
        Writeln("</head>")
        Writeln("<body>")
        Writeln("<div id=\"main\">")
        ParseMath(strings.Split(input,""),0).rend()
        Writeln("</div>")
        Writeln("</body>")
        Writeln("</html>")
    }
}