package main

import "fmt"
import "github.com/blue-granite/LearningGo/library/lib/sub"
import "github.com/blue-granite/LearningGo/library/lib"
import "github.com/blue-granite/LearningGo/library"

func main() {
	fmt.Println("Cmd for lib")
	sub.LibSubFunc()
	lib.LibFunc()

	//This function is at the root level of the module,
	//its package name is considered to be the module name,
	//hence imported via the module name.
	library.RootFunc()
}
