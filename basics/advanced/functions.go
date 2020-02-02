// functions
package advanced

import (
	"fmt"
)

func Functions() {
	fmt.Println("Function values may be used as function arguments and return values")
	functionAsValues()
	functionClosures()
}

func functionAsValues() {

	add := func(x, y float64) float64 {
		return x + y
	}

	multi := func(x, y float64) float64 {
		return x * y
	}

	a := funcThatTakeFuncAsArg(add)
	m := funcThatTakeFuncAsArg(multi)
	fmt.Println("a and m ", a, m)

	s := funcThatReturnsFuncAsArg(5)(3)
	fmt.Println("s ", s)

}

func funcThatTakeFuncAsArg(f func(float64, float64) float64) float64 {
	return f(3, 4)
}

func funcThatReturnsFuncAsArg(f float64) func(float64) float64 {
	return func(x float64) float64 {
		return x - f
	}
}

func functionClosures() {
	fmt.Println("Go functions may be closures, they can reference variables from outside its body.")

	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
}

func newCounter() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}
