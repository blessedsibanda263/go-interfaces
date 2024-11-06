package main

import "fmt"

type IntA interface {
	foo()
}

type IntB interface {
	bar()
}

type IntC interface {
	IntA
	IntB
}

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

func processA(s IntA) {
	fmt.Printf("%T\n", s)
}

// Satisfying IntA
func (varC c) foo() {
	fmt.Println("Foo Processing", varC)
}

// Satisfying IntB
func (varC c) bar() {
	fmt.Println("Bar Processing", varC)
}

// using an anonymouse structure inside another structure
type compose struct {
	field1 int
	a
}

// Different structures can have methods with the same name
func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) A() {
	fmt.Println("Function A() for B")
}

func main() {
	var iC c = c{A: a{XX: 120, YY: 12}, B: b{AA: "-12", XX: -12}}
	iC.A.A()
	iC.B.A()

	iComp := compose{field1: 123, a: a{XX: 120, YY: 12}}
	fmt.Println(iComp.XX, iComp.YY, iComp.field1)

	iC.bar()
	processA(iC)
}
