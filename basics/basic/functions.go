package basic

import (
	"fmt"
)

func Functions() {
	FuncNoParam()
	FuncTwoParam(1, 2)
	FuncParamShareType(3, 4)
	fmt.Printf("\t%s\n", FuncNoParamReturnType())
	fmt.Printf("\t%d\n", FuncTwoParamReturnType(5, 6))

	a, b := FuncMultipleReturnValues(7, 8)
	fmt.Printf("\tReturn values [%s] and [%d]\n", a, b)

	c, d := FuncNamedReturnValues(9, 10)
	fmt.Printf("\tNamed return values [%d] and [%d]\n", c, d)
}

func FuncNoParam() {
	fmt.Println("A function can take zero or more arguments, in this case none")
}

func FuncTwoParam(x int, y int) {
	fmt.Printf("A function can take zero or more arguments, in this case: %d %d\n", x, y)
}

func FuncParamShareType(x, y int) {
	fmt.Printf("A function can take zero or more arguments, in this case they are of the same type: %d %d\n", x, y)
}

func FuncNoParamReturnType() string {
	fmt.Println("A function can return a value")
	return "i'm returning a string"
}

func FuncTwoParamReturnType(x int, y int) int {
	fmt.Println("A function can return a value derived from params")
	return x + y
}

func FuncMultipleReturnValues(x int, y int) (string, int) {
	fmt.Println("A function can have multiple return values")
	return "added", x + y
}

func FuncNamedReturnValues(x int, y int) (xx, yy int) {
	fmt.Println("A function can have named return values")
	xx = x * x // no colon needed as xx already defined as return value in signature
	yy = y * y // no colon needed as yy already defined as return value in signature
	return     //will return named variables
}

//named return values
//multi retuen values
