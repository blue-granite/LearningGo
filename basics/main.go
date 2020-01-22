package main

import "fmt"
import "github.com/blue-granite/LearningGo/basics/basic"

func main() {
	fmt.Println("Basics")
	NoMixPkgInDir()
	basic.Packages()
	basic.Imports()
	basic.Exported()
	basic.FuncNone()
	basic.FuncTwo(1, 2)
	basic.FuncParamShareType(3, 4)
}
