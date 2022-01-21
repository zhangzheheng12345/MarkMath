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
	x.Rend()
}

func NewWhole(str string, li []Element) Whole {
	return Whole{str, li}
}

func (whole Whole) Rend() {
	li := strings.Split(whole.template, "##########")
	Writeln(li[0])
	for _, element := range whole.elements {
		element.Rend()
	}
	Writeln(li[1])
}

func NewLine(li []Element) Line {
	return Line{li}
}

func (line Line) Rend() {
	Writeln("<div>")
	for _, element := range line.elements {
		element.Rend()
	}
	Writeln("\n</div>")
}

func NewText(str string) Text {
	return Text{str}
}

func (text Text) Rend() {
	Write(text.value)
}

func NewNum(str string, next Element) Num {
	return Num{str, next}
}

func NewAlpha(str string, next Element) Alpha {
	return Alpha{str, next}
}

func (alpha Alpha) Rend() {
	Write(alpha.value)
	if alpha.next != nil {
		alpha.next.Rend()
	}
}

func (num Num) Rend() {
	Write(num.value)
	if num.next != nil {
		num.next.Rend()
	}
}

func NewPlus(next Element) Plus {
	return Plus{next}
}

func (plus Plus) Rend() {
	Write("+")
	if plus.next != nil {
		plus.next.Rend()
	}
}

func NewMinus(next Element) Minus {
	return Minus{next}
}

func (minus Minus) Rend() {
	Write("-")
	if minus.next != nil {
		minus.next.Rend()
	}
}

func NewMulti(next Element) Multi {
	return Multi{next}
}

func (multi Multi) Rend() {
	Writeln("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"14\" height=\"14\">")
	Writeln("<g transform=\"translate(3,6)\">")
	Writeln("<rect transform=\"rotate(45,4,1)\" width=\"8\" height=\"2\"/>")
	Writeln("<rect transform=\"rotate(135,4,1)\" width=\"8\" height=\"2\"/>")
	Writeln("</g>")
	Writeln("</svg>")
	if multi.next != nil {
		multi.next.Rend()
	}
}

func NewDivide(next Element) Divide {
	return Divide{next}
}

func (divide Divide) Rend() {
	Writeln("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"14\" height=\"14\">")
	Writeln("<circle cx=\"7\" cy=\"3\" r=\"1\"/>")
	Writeln("<rect transform=\"translate(2,6)\" width=\"10\" height=\"2\"/>")
	Writeln("<circle cx=\"7\" cy=\"11\" r=\"1\"/>")
	Writeln("</svg>")
	if divide.next != nil {
		divide.next.Rend()
	}
}

func NewEqual(next Element) Equal {
	return Equal{next}
}

func (equal Equal) Rend() {
	Write("=")
	if equal.next != nil {
		equal.next.Rend()
	}
}

func NewBigger(next Element) Bigger {
	return Bigger{next}
}

func (bigger Bigger) Rend() {
	Write(">")
	if bigger.next != nil {
		bigger.next.Rend()
	}
}

func NewSmaller(next Element) Smaller {
	return Smaller{next}
}

func (smaller Smaller) Rend() {
	Write("<")
	if smaller.next != nil {
		smaller.next.Rend()
	}
}

func NewFrontBrace(next Element) FrontBrace {
	return FrontBrace{next}
}

func (fBrace FrontBrace) Rend() {
	Write("(")
	if fBrace.next != nil {
		fBrace.next.Rend()
	}
}

func NewBackBrace(next Element) BackBrace {
	return BackBrace{next}
}

func (bBrace BackBrace) Rend() {
	Write(")")
	if bBrace.next != nil {
		bBrace.next.Rend()
	}
}

func NewPlusMinus(next Element) PlusMinus {
	return PlusMinus{next}
}

func (plusMinus PlusMinus) Rend() {
	Writeln("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"13\" height=\"13\">")
	Writeln("<rect transform=\"translate(6,3)\" width=\"1\" height=\"7\"/>")
	Writeln("<rect transform=\"translate(3,6)\" width=\"7\" height=\"1\"/>")
	Writeln("<rect transform=\"translate(3,11)\" width=\"7\" height=\"1\"/>")
	Writeln("</svg>")
	if plusMinus.next != nil {
		plusMinus.next.Rend()
	}
}

func NewFraction(top Element, bottom Element, next Element) Fraction {
	return Fraction{top, bottom, next}
}

func (fraction Fraction) Rend() {
	Writeln("<div class=\"fraction\">")
	Writeln("<div>")
	fraction.top.Rend()
	Writeln("</div>")
	Writeln("<div class=\"fractionLine\">")
	Writeln("</div>")
	Writeln("<div>")
	fraction.bottom.Rend()
	Writeln("</div>")
	Writeln("</div>")
	if fraction.next != nil {
		fraction.next.Rend()
	}
}

func NewRoot(top Element, bottom Element, next Element) Root {
	return Root{top, bottom, next}
}

func (root Root) Rend() {
	Writeln("<div class=\"root\">")
	Writeln("<div class=\"rootFront\">")
	Writeln("<div>")
	root.rootNum.Rend()
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
	root.rootBottom.Rend()
	Writeln("</div>")
	Writeln("</div>")
	Writeln("</div>")
	if root.next != nil {
		root.next.Rend()
	}
}

func NewPower(point Element, next Element) Power {
	return Power{point, next}
}

func (power Power) Rend() {
	Writeln("<sup>")
	power.point.Rend()
	Writeln("</sup>")
	if power.next != nil {
		power.next.Rend()
	}
}
func NewSubscript(value Element, next Element) Subscript {
	return Subscript{value, next}
}

func (subscript Subscript) Rend() {
	Writeln("<sub>")
	subscript.value.Rend()
	Writeln("</sub>")
	if subscript.next != nil {
		subscript.next.Rend()
	}
}
