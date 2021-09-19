package main

import (
    "fmt"
    "strings"
    "elements"
)

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
            for IsSingleNum(str[endIndex]) {
                endIndex++
            }
            return NewNum(strings.Join(str[index:endIndex],""),ParseMath(str,endIndex))
        }
    }
    return nil
}