package elements

import (
	"fmt"
	"strings"
)

func Write(str string) {
	fmt.Print(str)
}

func Writeln(str string) {
	fmt.Println(str)
}

func Rend(x Element) {
    x.rend()
}

func NewWhole(str string,li []Element) Whole {
    return Whole{str,li}
}

func (whole Whole) rend() {
	li := strings.Split(whole.template, "##########")
	Writeln(li[0])
	for _, element := range whole.elements {
		element.rend()
	}
	Writeln(li[1])
}

func NewLine(li []Element) Line {
	return Line{li}
}

func (line Line) rend() {
	Writeln("<div>")
	for _, element := range line.elements {
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

func NewNum(str string, next Element) Num {
	return Num{str, next}
}

func NewAlpha(str string, next Element) Alpha {
	return Alpha{str, next}
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
	Writeln("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"13\" height=\"13\">")
	Writeln("<rect transform=\"translate(6,3)\" width=\"1\" height=\"7\"/>")
	Writeln("<rect transform=\"translate(3,6)\" width=\"7\" height=\"1\"/>")
	Writeln("<rect transform=\"translate(3,11)\" width=\"7\" height=\"1\"/>")
	Writeln("</svg>")
	if plusMinus.next != nil {
		plusMinus.next.rend()
	}
}

func NewFraction(top Element, bottom Element, next Element) Fraction {
	return Fraction{top, bottom, next}
}

func (fraction Fraction) rend() {
	Writeln("<div class=\"fraction\">")
	Writeln("<div>")
	fraction.top.rend()
	Writeln("</div>")
	Writeln("<div class=\"fractionLine\">")
	Writeln("</div>")
	Writeln("<div>")
	fraction.bottom.rend()
	Writeln("</div>")
	Writeln("</div>")
	if fraction.next != nil {
		fraction.next.rend()
	}
}

func NewRoot(top Element, bottom Element, next Element) Root {
	return Root{top, bottom, next}
}

func (root Root) rend() {
	Writeln("<div class=\"root\">")
	Writeln("<div class=\"rootFront\">")
	Writeln("<div>")
	root.rootNum.rend()
	Writeln("</div>")
	Writeln("<svg width=\"4px\" height=\"3px\" xmlns=\"http://www.w3.org/2000/svg\">")
	Writeln("	<path d=\"M0 2 L2 0 L4 3\" stroke=\"black\"/>")
	Writeln("</svg>")
	Writeln("</div>")
	Writeln("<div class=\"rootMid\">")
	Writeln("</div>")
	Writeln("<div class=\"rootBack\">")
	Writeln("<div class=\"rootLine\">")
	Writeln("</div>")
	Writeln("<div>")
	root.rootBottom.rend()
	Writeln("</div>")
	Writeln("</div>")
	Writeln("</div>")
	if root.next != nil {
		root.next.rend()
	}
}

func NewPower(point Element, next Element) Power {
	return Power{point, next}
}

func (power Power) rend() {
	Writeln("<sup>")
	power.point.rend()
	Writeln("</sup>")
	if power.next != nil {
		power.next.rend()
	}
}
func NewSubscript(value Element, next Element) Subscript {
	return Subscript{value, next}
}

func (subscript Subscript) rend() {
	Writeln("<sub>")
	subscript.value.rend()
	Writeln("</sub>")
	if subscript.next != nil {
		subscript.next.rend()
	}
}