package elements

type Element interface {
	rend()
}

type Whole struct {
	template string
	elements []Element
}

type Line struct {
	elements []Element
}

type Text struct {
	value string
}

type Num struct {
	value string
	next  Element
}

type Alpha struct {
	value string
	next  Element
}

type Plus struct {
	next Element
}

type Minus struct {
	next Element
}

type Multi struct {
	next Element
}

type Divide struct {
	next Element
}

type Equal struct {
	next Element
}

type Bigger struct {
	next Element
}

type Smaller struct {
	next Element
}

type FrontBrace struct {
	next Element
}

type BackBrace struct {
	next Element
}

type PlusMinus struct {
	next Element
}

type Fraction struct {
	top    Element
	bottom Element
	next   Element
}
