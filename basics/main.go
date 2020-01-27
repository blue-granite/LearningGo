package main

import "fmt"
import "github.com/blue-granite/LearningGo/basics/basic"
import "github.com/blue-granite/LearningGo/basics/flow"

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
}
