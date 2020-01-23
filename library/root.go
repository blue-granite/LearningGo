package library

import (
	"fmt"
)

func RootFunc() {
	fmt.Println("Func at root of module, this file is at the root level of the module, its package name is considered to be the module name, hence imported via the module name. ")
}
