package parser

import (
	"strings"

	"github.com/zhangzheheng12345/MarkMath/parser/elements"
)

const template = `<!DOCTYPE html>
<html>
    <head>
        <style>
            div {
                display:flex;
                align-items:center;
                padding-left:1px;
                padding-right:1px;
            }
            div.fraction {
                flex-direction:column;
            }
            div.fractionLine {
                width:100%;
                height:1px;
                background-color:black;
            }
            div .root {
                display:flex;
                align-items:center;
                margin-left:1px;
                margin-right:1px;
            }
            div .rootFront {
                flex-direction:column;
                align-items:flex-end;
                justify-content:flex-end;
                padding:0px;
                border-right:0px;
            }
            div .rootMid {
                background-color:black;
                width:2px;
                height:100%;
            }
            div .rootLine {
                background-color:black;
                width:100%;
                height:1px;
            }
            div .rootBack {
                flex-direction:column;
                align-items:flex-end;
                padding:0px;
                border-left:0px;
            }
            body {
                padding:15px;
            }
        </style>
    </head>
    <body>
        <!-- MarkMath Mark : ten ' # ' are where will be filled. -->
        ##########
    </body>
</html>`

func Rend(e elements.Element) {
	e.Rend()
}

func ParseWhole(strs []string) elements.Element {
	var result []elements.Element
	for _, str := range strs {
		result = append(result, ParseLine(strings.Split(str, "")))
	}
	return elements.NewWhole(template, result)
}

func ParseLine(str []string) elements.Element {
	var index int = 0
	var result []elements.Element
	var begin int = 0
	for ; index < len(str); index++ {
		if str[index] == "$" && len(str)-index >= 2 {
			/* no space between $[ ] */
			if str[index+1] == "[" {
				result = append(result, elements.NewText(strings.Join(str[begin:index], "")))
				var count int = 1
				index += 2
				begin = index
				for ; count != 0; index++ {
					if str[index] == "[" {
						count++
					} else if str[index] == "]" {
						count--
					}
				}
				result = append(result, ParseMath(str[begin:index-1], 0))
				begin = index
			}
		}
	}
	if index-begin > 1 {
		result = append(result, elements.NewText(strings.Join(str[begin:index], "")))
	}
	return elements.NewLine(result)
}

func IsSingleNum(str string) bool {
	return len(str) == 1 && []byte(str)[0] >= 48 && []byte(str)[0] <= 57
}

func ParseMath(str []string, index int) elements.Element {
	if index < len(str) {
		if str[index] == "+" {
			if index+1 < len(str) && str[index+1] == "-" {
				return elements.NewPlusMinus(ParseMath(str, index+2))
			} else {
				return elements.NewPlus(ParseMath(str, index+1))
			}
		} else if str[index] == "-" {
			return elements.NewMinus(ParseMath(str, index+1))
		} else if str[index] == "*" {
			return elements.NewMulti(ParseMath(str, index+1))
		} else if str[index] == "/" {
			if index+1 < len(str) && str[index+1] == "[" {
				index += 2
				count := 0
				result := []string{}
				for index < len(str) && !(count == 0 && str[index] == ",") {
					result = append(result, str[index])
					if str[index] == "[" {
						count++
					} else if str[index] == "]" {
						count--
					}
					index++
				}
				top := ParseMath(result, 0)
				index++
				count = 0
				result = []string{}
				for index < len(str) && !(count == 0 && str[index] == "]") {
					result = append(result, str[index])
					if str[index] == "[" {
						count++
					} else if str[index] == "]" {
						count--
					}
					index++
				}
				bottom := ParseMath(result, 0)
				return elements.NewFraction(top, bottom, ParseMath(str, index+1))
			} else {
				return elements.NewDivide(ParseMath(str, index+1))
			}
		} else if str[index] == "\\" {
			index += 2
			count := 0
			result := []string{}
			for index < len(str) && !(count == 0 && str[index] == ",") {
				result = append(result, str[index])
				if str[index] == "[" {
					count++
				} else if str[index] == "]" {
					count--
				}
				index++
			}
			top := ParseMath(result, 0)
			index++
			count = 0
			result = []string{}
			for index < len(str) && !(count == 0 && str[index] == "]") {
				result = append(result, str[index])
				if str[index] == "[" {
					count++
				} else if str[index] == "]" {
					count--
				}
				index++
			}
			bottom := ParseMath(result, 0)
			return elements.NewRoot(top, bottom, ParseMath(str, index+1))
		} else if str[index] == "=" {
			return elements.NewEqual(ParseMath(str, index+1))
		} else if str[index] == ">" {
			return elements.NewBigger(ParseMath(str, index+1))
		} else if str[index] == "<" {
			return elements.NewSmaller(ParseMath(str, index+1))
		} else if str[index] == "(" {
			return elements.NewFrontBrace(ParseMath(str, index+1))
		} else if str[index] == ")" {
			return elements.NewBackBrace(ParseMath(str, index+1))
		} else if IsSingleNum(str[index]) {
			endIndex := index
			for endIndex < len(str) && (IsSingleNum(str[endIndex]) || str[endIndex] == ".") {
				endIndex++
			}
			return elements.NewNum(strings.Join(str[index:endIndex], ""), ParseMath(str, endIndex))
		} else if str[index] == ";" {
			index++
			if index < len(str) {
				if str[index] == "a" {
					return elements.NewAlpha("α", ParseMath(str, index+1))
				} else if str[index] == "b" {
					return elements.NewAlpha("β", ParseMath(str, index+1))
				} else if str[index] == "d" {
					return elements.NewAlpha("Δ", ParseMath(str, index+1))
				} else if str[index] == "l" {
					return elements.NewAlpha("λ", ParseMath(str, index+1))
				} else if str[index] == "t" {
					return elements.NewAlpha("θ", ParseMath(str, index+1))
				} else if str[index] == "p" {
					return elements.NewAlpha("π", ParseMath(str, index+1))
				} else if str[index] == "g" {
					return elements.NewAlpha("γ", ParseMath(str, index+1))
				}
			}
		} else if str[index] == "^" {
			index += 2
			count := 0
			result := []string{}
			for index < len(str) && !(count == 0) {
				result = append(result, str[index])
				if str[index] == "[" {
					count++
				} else if str[index] == "]" {
					count--
				}
				index++
			}
			point := ParseMath(result, 0)
			return elements.NewPower(point, ParseMath(str, index+1))
		} else if str[index] == "." {
			index += 2
			count := 0
			result := []string{}
			for index < len(str) && !(count == 0) {
				result = append(result, str[index])
				if str[index] == "[" {
					count++
				} else if str[index] == "]" {
					count--
				}
				index++
			}
			point := ParseMath(result, 0)
			return elements.NewSubscript(point, ParseMath(str, index+1))
		} else {
			return elements.NewAlpha(str[index], ParseMath(str, index+1))
		}
	}
	return nil
}
