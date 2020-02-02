package methodsandinterfaces

import (
	"fmt"
	"math"
)

//--------------------------

type Foo struct {
	X float64
	Y float64
}

func (f Foo) Abs() float64 {
	return math.Sqrt(f.X*f.X + f.Y*f.Y)
}

func (f *Foo) PointerReciever(s string) {
	fmt.Println("PointerReciever", s, f)
}

func (f Foo) ValueReciever(s string) {
	fmt.Println("ValueReciever", s, f)
}

//---------------------------

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//--------------------------

func Methods() {
	fmt.Println("You can define methods on types.")
	fmt.Println("A method is a function with a special receiver argument.")
	fmt.Println("\tfunc (f Foo) Abs() float64 ")

	f := Foo{3, 4}
	fmt.Println(f.Abs())

	fmt.Println("You can define methods on non struct types.")

	var m MyFloat = -7
	fmt.Println(m.Abs())

	fmt.Println("You can only declare a method with a receiver whose type is defined in the same package as the method.")

	fmt.Println("You can declare methods with pointer receivers.")
	fPointer := &f

	f.PointerReciever("invoked on value, intepreted as (&f)")
	f.ValueReciever("invoked on value")

	fPointer.PointerReciever("invoked on pointer")
	fPointer.ValueReciever("invoked on pointer, intepreted as (*fPointer)")

	fmt.Println("There are two reasons to use a pointer receiver.")
	fmt.Println("\t1 method can modify the value that its receiver points to.")
	fmt.Println("\t2 to avoid copying the value on each method call.")
	fmt.Println("In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.")

}
