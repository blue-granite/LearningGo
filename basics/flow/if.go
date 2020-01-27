package flow

import (
	"fmt"
)

func If() {
	x := 3

	if x > 2 {
		fmt.Println("If with ()")
	}

	if x > 2 {
		fmt.Println("If without ()")
	}

	if x < 2 {
		fmt.Println("x < 2")
	} else {
		fmt.Println("If-else")
	}

	if x < 2 {
		fmt.Println("x < 2")
	} else if x > 0 {
		fmt.Println("If-else-if")
	}

	if y := 3; y > 2 {
		fmt.Println("If starting with a short statement")
	}
}
