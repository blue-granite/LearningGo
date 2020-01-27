package flow

import (
	"fmt"
	"runtime"
	"time"
)

func Switch() {
	SwitchOn()
	SwitchOnWithStatement()
	SwitchNoCondition()
}

func SwitchOn() {

	os := runtime.GOOS

	fmt.Println("Go's switch cases need not be constants, and the values involved need not be integers.")
	fmt.Println("Go's switch cases are evaluated from top to bottom.")
	fmt.Println("Go's switch only runs the matched case, no fall-through, no need for break.")

	fmt.Printf("%s %T.\n", os, os)

	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

}

func SwitchOnWithStatement() {

	fmt.Println("Go's switch starting with a short statement.")

	switch max := runtime.NumCPU(); max {
	case 1:
		fmt.Println("Only one proc")
	case 2:
		fmt.Println("Only two proc")
	default:
		fmt.Printf("%d procs.\n", max)
	}

}

func SwitchNoCondition() {

	fmt.Println("Go's switch doesn't need a condition, like a long if-else-if... block.")

	t := time.Now()
	switch /* see nothing here */ {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
