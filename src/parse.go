package main

import (
    "fmt"
    "strings"
    "elements"
)

func ParseWhole(strs []string) Element {
    var result []elements.Element
    for _,str := range strs {
        result = append(result,ParseLine(strings.Split(str,"")))b
    }
    return elements.NewWhole(result)
}

func ParseLine(str []string) Element {
    var index int = 0
    var result []elements.Element
    var begin int = 0
    for ;index < len(str);index++ {
        if str[index] == "$" && len(str) - index >= 2{
            /* no space between $[ ] */
            if str[index + 1] == "[" { 
                result = append(result,elements.NewText(strings.Join(str[begin:index],"")))
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
    return elements.NewLine(result)
}

func IsSingleNum(str string) bool {
    return len(str) == 1 && []byte(str)[0] >= 48 && []byte(str)[0] <= 57
}

func IsSingleAlpha(str string) bool {
    return len(str) == 1 && []byte(str)[0] >= 65 && []byte(str)[0] <= 122
}

func ParseMath(str []string,index int) Element {
    if index < len(str) {
        if str[index] == "+" {
            if index + 1 < len(str) && str[index + 1] == "-" {
                return elements.NewPlusMinus(ParseMath(str,index + 2))
            } else {
                return elements.NewPlus(ParseMath(str,index + 1))
            }
        } else if str[index] == "-" {
            return elements.NewMinus(ParseMath(str,index + 1))
        } else if str[index] == "*" {
            return elements.NewMulti(ParseMath(str,index + 1))
        } else if str[index] == "/" {
            return elements.NewDivide(ParseMath(str,index + 1))
        } else if str[index] == "=" {
            return elements.NewEqual(ParseMath(str,index + 1))
        } else if IsSingleNum(str[index]) {
            endIndex := index
            for endIndex < len(str) && IsSingleNum(str[endIndex]) {
                endIndex++
            }
            return elements.NewNum(strings.Join(str[index:endIndex],""),ParseMath(str,endIndex))
        } else if IsSingleAlpha(str[index]) {
            return elements.NewAlpha(str[index],ParseMath(str,index+1))
        } else {
            return ParseMath(str,index + 1)
        }
    }
    return nil
}