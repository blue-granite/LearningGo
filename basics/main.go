package main

import "fmt"
import "github.com/blue-granite/LearningGo/basics/basic"
import "github.com/blue-granite/LearningGo/basics/flow"
import "github.com/blue-granite/LearningGo/basics/advanced"
import "github.com/blue-granite/LearningGo/basics/methodsandinterfaces"
import "github.com/blue-granite/LearningGo/basics/goroutines"

const LINE = "---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------"

func line() {

	fmt.Println(LINE)
}
func main() {
	fmt.Println("Basics")
	NoMixPkgInDir()
	line()
	basic.Packages()
	line()
	basic.Imports()
	line()
	basic.Exported()
	line()
	basic.Functions()
	line()
	basic.Variables()
	line()
	basic.Constants()
	line()
	basic.Types()
	line()
	flow.For()
	line()
	flow.If()
	line()
	flow.Switch()
	line()
	flow.Defer()
	line()
	advanced.Pointers()
	line()
	advanced.Structs()
	line()
	advanced.Arrays()
	line()
	advanced.Slices()
	line()
	advanced.SlicesOfSlices()
	line()
	advanced.SlicesMake()
	line()
	advanced.SlicesAppend()
	line()
	advanced.Ranges()
	line()
	advanced.Maps()
	line()
	advanced.Functions()
	line()
	methodsandinterfaces.Methods()
	line()
	methodsandinterfaces.Interfaces()
	line()
	goroutines.GoRoutines()
	line()
	goroutines.Mutexes()
}
