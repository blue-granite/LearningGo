package main

import "fmt"
import "github.com/blue-granite/LearningGo/library/lib/sub"
import "github.com/blue-granite/LearningGo/library/lib"
import "github.com/blue-granite/LearningGo/library"

func main() {
	fmt.Println("Modules")
	library.RootFunc()
	lib.LibFunc()
	sub.LibSubFunc()
}
