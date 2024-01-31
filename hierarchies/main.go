package main

import "fmt"

type foo float32
type intfoo int

func (f foo) String() string {
	return "Foo " + fmt.Sprintf("%f", f)
}

type  foofoo struct {
	f, f2 foo
}

type foofoobar foofoo

type ifooifoo struct {
	iff, iff2 intfoo
}

type ifooifoo2 struct {
	iff, iff2 intfoo
}

type ifooifoo3 struct {
	iff3, iff4 intfoo
}

type bar foo
func (b bar) String() string {
	return "Bar " + fmt.Sprintf("%f", b)
}

func main() {
	var f foo = 3.0
	var b bar = 4.0
	fmt.Println("1f: ", f, "b: ", b)
	b = bar(f) // cast always required - no up/downcasting
	fmt.Println("2f: ", f, "b: ", b)
	b = 44
	f = foo(b)
	fmt.Println("3f: ", f, "b: ", b)

	fmt.Println("------------")
	var ifoo intfoo = 444
	fmt.Println("ifoo", ifoo)
	ifoo = intfoo(f)
	fmt.Println("2ifoo", ifoo)
	
	fmt.Println("------------")
	var fuf foofoo = foofoo{1.1,2.2}
	
	//fuf = fuf(f) - compilation error
	//f = foo(fuf) - compilation error
	fmt.Println("fuf: ", fuf) //  works but does not call String from foo but from floats
	
	var ifuf ifooifoo = ifooifoo{1,2}
	fmt.Println("ifuf: ", ifuf)
	
	//ifuf = ifooifoo(fuf) - compilation error
	//fuf = foofoo(ifuf) - compilation error
	
	var fufbar foofoobar = foofoobar(fuf)
	fmt.Println("fufbar from fuf", fufbar)
	
	fufbar.f = 100
	fufbar.f2 = 200
	fuf = foofoo(fufbar)
	fmt.Println("fuf from fufbar", fuf)
	
	fmt.Println("------------")
	var ifuf2 ifooifoo2 = ifooifoo2{10,20}
	fmt.Println("ifuf2: ", ifuf2)
	
	ifuf2  = ifooifoo2(ifuf) // surprise - this actually works!
	// var ifuf3 ifooifoo3 = ifooifoo3(ifuf) - but this is compile error - field names different
	fmt.Println("ifuf2 from ifuf: ", ifuf2)


}