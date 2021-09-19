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
    /* Use svg later */
    Write("*")
    if multi.next != nil {
        multi.next.rend()
    }
}

func NewDivide(next Element) Divide {
    return Divide{next}
}

func (divide Divide) rend() {
    /* Use svg later */
    Write("/")
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