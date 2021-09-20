package main

import (
    "fmt"
    "strings"
    "elements"
)

func ParseLine(str []string) Element{
    var index int = 0
    var result []Element
    var begin int = 0
    for ;index < len(str);index++ {
        if str[index] == "$" && len(str) - index >= 2{
            /* no space between &[ ] */
            if str[index + 1] == "[" { 
                result = append(result,NewText(strings.Join(str[begin:index],"")))
                var count int = 1
                index += 2
                begin = index
                for ;count != 0;index++ {
                    if str[index] == "[" {
                        count++
                    } else if str[index] == "]" {
                        count--
                    }
                }
                result = append(result,ParseMath(str[begin:index - 1],0))
                begin = index
            }
        }
    }
    if index - begin > 1{
        result = append(result,NewText(strings.Join(str[begin:index],"")))
    }
    return NewLine(result)
}

func IsSingleNum(str string) bool {
    return len(str) == 1 && []byte(str)[0] >= 48 && []byte(str)[0] <= 57
}

func ParseMath(str []string,index int) Element{
    if len(str) - index == 1 {
        if str[index] == "+" {
            return NewPlus(nil)
        } else if str[index] == "-" {
            return NewMinus(nil)
        } else if str[index] == "*" {
            return NewMulti(nil)
        } else if str[index] == "/" {
            return NewDivide(nil)
        } else if str[index] == "=" {
            return NewEqual(nil)
        } else if IsSingleNum(str[index]) {
            return NewNum(str[index],nil)
        }
    } else {
        if str[index] == "+" {
            return NewPlus(ParseMath(str,index + 1))
        } else if str[index] == "-" {
            return NewMinus(ParseMath(str,index + 1))
        } else if str[index] == "*" {
            return NewMulti(ParseMath(str,index + 1))
        } else if str[index] == "/" {
            return NewDivide(ParseMath(str,index + 1))
        } else if str[index] == "=" {
            return NewEqual(ParseMath(str,index + 1))
        } else if IsSingleNum(str[index]) {
            endIndex := index
            for IsSingleNum(str[endIndex]) && endIndex < len(str) {
                endIndex++
            }
            return NewNum(strings.Join(str[index:endIndex],""),ParseMath(str,endIndex))
        }
    }
    return nil
}