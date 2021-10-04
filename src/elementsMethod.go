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

func NewWhole(li []Element) Whole {
    return Whole{li}
}

func (whole Whole) rend() {
    for _,element := range whole.elements {
        element.rend()
    }
}

func NewLine(li []Element) Line {
    return Line{li}
}

func (line Line) rend() {
    Writeln("<div>")
    for _,element := range line.elements {
        element.rend()
    }
    Writeln("\n</div>")
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

func NewAlpha(str string,next Element) Alpha {
    return Alpha{str,next}
}

func (alpha Alpha) rend() {
    Write(alpha.value)
    alpha.next.rend()
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
    Writeln("<g transform=\"translate(3,6)\">")
    Writeln("<rect transform=\"rotate(45,4,1)\" width=\"8\" height=\"2\"/>")
    Writeln("<rect transform=\"rotate(135,4,1)\" width=\"8\" height=\"2\"/>")
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

func NewBigger(next Element) Bigger {
    return Bigger{next}
}

func (bigger Bigger) rend() {
    Write(">")
    if bigger.next != nil {
        bigger.next.rend()
    }
}

func NewSmaller(next Element) Smaller {
    return Smaller{next}
}

func (smaller Smaller) rend() {
    Write("<")
    if smaller.next != nil {
        smaller.next.rend()
    }
}

func NewFrontBrace(next Element) FrontBrace {
    return FrontBrace{next}
}

func (fBrace FrontBrace) rend() {
    Write("(")
    if fBrace.next != nil {
        fBrace.next.rend()
    }
}

func NewBackBrace(next Element) BackBrace {
    return BackBrace{next}
}

func (bBrace BackBrace) rend() {
    Write(")")
    if bBrace.next != nil {
        bBrace.next.rend()
    }
}

func NewPlusMinus(next Element) PlusMinus {
    return PlusMinus{next}
}

func (plusMinus PlusMinus) rend() {
    Writeln("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"14\" height=\"14\">")
    Writeln("<rect transform=\"translate(6,2)\" width=\"2\" height=\"8\"/>")
    Writeln("<rect transform=\"translate(2,6)\" width=\"10\" height=\"2\"/>")
    Writeln("<rect transform=\"translate(2,11)\" width=\"10\" height=\"2\"/>")
    Writeln("</svg>")
    if plusMinus.next != nil {
        plusMinus.next.rend()
    }
}