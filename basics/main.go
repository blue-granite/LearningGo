package main

import "fmt"
import "github.com/blue-granite/LearningGo/basics/basic"
import "github.com/blue-granite/LearningGo/basics/flow"
import "github.com/blue-granite/LearningGo/basics/advanced"

func main() {
	fmt.Println("Basics")
	NoMixPkgInDir()
	basic.Packages()
	basic.Imports()
	basic.Exported()
	basic.Functions()
	basic.Variables()
	basic.Constants()
	basic.Types()
	flow.For()
	flow.If()
	flow.Switch()
	flow.Defer()
	advanced.Pointers()
	advanced.Structs()
}
