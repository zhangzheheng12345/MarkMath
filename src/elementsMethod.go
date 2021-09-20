package elements

import (
    "fmt"
)

func Write(str string) {
    fmt.Print(str)
}

func Writeln(str string) {
    fmt.Println(str)
}

func NewLine(li []Element) Line {
    return Line{li}
}

func (line Line) rend() {
    for _,element := line.elements {
        element.rend()
    }
    Write("\n")
}

func NewText(str string) Text {
    return Text{str}
}

func (text Text) rend() {
    Write(text.value)
}

func NewNum(str string,next Element) Num {
    return Num{str,next}
}

func (num Num) rend() {
    Write(num.value)
    if num.next != nil {
        num.next.rend()
    }
}

func NewPlus(next Element) Plus {
    return Plus{next}
}

func (plus Plus) rend() {
    Write("+")
    if plus.next != nil {
        plus.next.rend()
    }
}

func NewMinus(next Element) Minus {
    return Minus{next}
}

func (minus Minus) rend() {
    Write("-")
    if minus.next != nil {
        minus.next.rend()
    }
}

func NewMulti(next Element) Multi {
    return Multi{next}
}

func (multi Multi) rend() {
    Writeln("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"14\" height=\"14\">")
    Writeln("<g transform=\"translate(2,5)\">")
    Writeln("<rect transform=\"rotate(45,5,1)\" width=\"10\" height=\"2\"/>")
    Writeln("<rect transform=\"rotate(135,5,1)\" width=\"10\" height=\"2\"/>")
    Writeln("</g>")
    Writeln("</svg>")
    if multi.next != nil {
        multi.next.rend()
    }
}

func NewDivide(next Element) Divide {
    return Divide{next}
}

func (divide Divide) rend() {
    Writeln("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"14\" height=\"14\">")
    Writeln("<circle cx=\"7\" cy=\"3\" r=\"1\"/>")
    Writeln("<rect transform=\"translate(2,6)\" width=\"10\" height=\"2\"/>")
    Writeln("<circle cx=\"7\" cy=\"11\" r=\"1\"/>")
    Writeln("</svg>")
    if divide.next != nil {
        divide.next.rend()
    }
}

func NewEqual(next Element) Equal {
    return Equal{next}
}

func (equal Equal) rend() {
    Write("=")
    if equal.next != nil {
        equal.next.rend()
    }
}