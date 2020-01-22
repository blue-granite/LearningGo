package basic

import (
	"fmt"
)

func Exported() {
	fmt.Println("This func is EXPORTED and visible to other packages.")
}

func notExported() {
	fmt.Println("This func is NOT exported and visible to other packages.")
}
