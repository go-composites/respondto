package a

import (
	"reflect"

	"b"
)

// --- a local stand-in for go-composites' Object/Base package ---------------
// RespondTo(object, name) is the package-level reflective helper the
// composition style uses; the analyzer resolves the first argument's type.

func RespondTo(object interface{}, methodName string) bool {
	return reflect.ValueOf(object).MethodByName(methodName).IsValid()
}

// --- types under test -------------------------------------------------------

// Greeter is a named interface with a known method set.
type Greeter interface {
	Hello() string
	Bye() string
}

// Base contributes Ping via embedding.
type Base interface{ Ping() string }

// Derived gets Ping through embedding Base, plus its own Bar.
type Derived interface {
	Base
	Bar() int
}

type widget struct{}

func (widget) Hello() string { return "hi" }
func (widget) Bye() string   { return "bye" }
func (widget) RespondTo(methodName string) bool {
	return RespondTo(widget{}, methodName)
}

// gadget carries a pointer method to exercise pointer-receiver resolution.
type gadget struct{}

func (*gadget) Run() string { return "go" }

// --- exercises --------------------------------------------------------------

func packageLevel(g Greeter, d Derived) {
	_ = RespondTo(g, "Hello") // ok: Greeter has Hello
	_ = RespondTo(g, "Nope")  // want `RespondTo: Greeter has no method "Nope"`
	_ = RespondTo(d, "Ping")  // ok: Ping is promoted from embedded Base
	_ = RespondTo(d, "Gone")  // want `RespondTo: Derived has no method "Gone"`
}

func methodForm() {
	w := widget{}
	_ = w.RespondTo("Hello") // ok: widget has Hello
	_ = w.RespondTo("Typo")  // want `RespondTo: widget has no method "Typo"`
}

func qualified(g Greeter) {
	_ = b.RespondTo(g, "Bye")  // ok: qualified package-level call, Bye exists
	_ = b.RespondTo(g, "Nada") // want `RespondTo: Greeter has no method "Nada"`
}

func pointers(g *gadget, w widget) {
	_ = RespondTo(g, "Run")   // ok: pointer to named struct, Run is a ptr method
	_ = RespondTo(g, "Walk")  // want `RespondTo: gadget has no method "Walk"`
	_ = RespondTo(w, "Hello") // ok: named struct value, Hello exists
	_ = RespondTo(w, "Halt")  // want `RespondTo: widget has no method "Halt"`
}

func dynamic(g Greeter, name string) {
	_ = RespondTo(g, name)         // ok: non-literal argument, ignored
	_ = RespondTo(g, "He"+"llo")   // ok: folds to a real method name
	var any interface{} = g
	_ = RespondTo(any, "Whatever") // ok: empty interface, no method set to judge
	var p *int
	_ = RespondTo(p, "Nope")       // ok: pointer to unnamed type, no method set
}
