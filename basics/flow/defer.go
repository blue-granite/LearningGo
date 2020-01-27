package flow

import (
	"fmt"
)

func Defer() {

	deferNoArgs()
	deferArgs()
	deferClosure()
}

func deferNoArgs() {
	fmt.Println("Defer")
	defer aFunc()
	defer bFunc()
	cFunc()
}

func deferArgs() {

	fmt.Println("Defer args evaluated immediatley")
	n := 0
	defer aFuncArg(n)
	n++
	defer bFuncArg(n)
	n++
	cFuncArg(n)
}

func deferClosure() {

	fmt.Println("Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it. ")
	defer func() {
		fmt.Println("deferred anonymous function ")
	} /*() make it a function call */ ()
	fmt.Println("last line of function ")
}

func aFunc() {
	fmt.Println("aaa")
}

func bFunc() {
	fmt.Println("bbb")
}

func cFunc() {
	fmt.Println("ccc")
}

func aFuncArg(x int) {
	fmt.Println("aaa", x)
}

func bFuncArg(x int) {
	fmt.Println("bbb", x)
}

func cFuncArg(x int) {
	fmt.Println("ccc", x)
}
