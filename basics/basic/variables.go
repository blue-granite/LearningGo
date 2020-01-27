package basic

import (
	"fmt"
)

func Variables() {

	var a, b, c bool
	fmt.Printf("var statement declares a list of variables, type is that end: %t %t %t\n", a, b, c)

	var i, j int = 1, 2
	fmt.Printf("var declaration can include initializers, one per variable: %d %d\n", i, j)

	s := "string"
	fmt.Printf("Short var declaration with implicit type: %s\n", s)

	t := 1.23
	fmt.Printf("Short var declaration with implicit type, type inference depends on precision: %v %T\n", t, t)

	u, v := "one", "two"
	fmt.Printf("Multiple short var declaration with implicit type: %s %s\n", u, v)

	w, x := "three", 3
	fmt.Printf("Multiple short var declaration with mixed implicit types: %s %d\n", w, x)

	var (
		m bool = false
		n int  = 4
	)
	fmt.Printf("var declaration can span multiple lines and types: %t %d\n", m, n)
}

const Foo = "Foo"
const (
	Bar         = 1.23
	Lar float64 = 42 //typed
)

func Constants() {
	fmt.Printf("const Foo: %v %T\n", Foo, Foo)
	fmt.Printf("const Bar: %v %T\n", Bar, Bar)
	fmt.Printf("const Lar: %v %T\n", Lar, Lar)
	fmt.Println("Numeric constants are high-precision values, an untyped constant takes the type needed by its context.")
	needFloat32(Bar)
	needFloat64(Bar)

}

func needFloat32(x float32) {
	fmt.Printf("const untyped, float32 asked for: %v %T\n", x, x)
}

func needFloat64(x float64) {
	fmt.Printf("const untyped, float64 asked for: %v %T\n", x, x)
}
