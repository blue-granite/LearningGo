package basic

import (
	"fmt"
)

func FuncNone() {
	fmt.Println("A function can take zero or more arguments, in this case none")

}

func FuncTwo(x int, y int) {
	fmt.Printf("A function can take zero or more arguments, in this case: %d %d\n", x, y)

}

func FuncParamShareType(x, y int) {
	fmt.Printf("A function can take zero or more arguments, in this case they are of the same type: %d %d\n", x, y)

}

//name return values
//multi retuen values
