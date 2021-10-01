package elements

import (
    "fmt"
)

type Element interface {
    rend()
}

type Whole struct {
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
    next Element
}

type Alpha struct {
    value string
    next Element
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

type PlusMinus struct {
    next Element
}