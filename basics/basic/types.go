package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

func Types() {
	BasicTypes()
	ZeroValues()
	TypeConversion()
}

func BasicTypes() {

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("Basic types: bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte (alias for uint8), rune (alias for int32, represents a Unicode code point), float32, float64, complex64, complex128")

}

func ZeroValues() {

	fmt.Println("Zero value for string is: \"\"")
	fmt.Println("Zero value for bool is: false)")
	fmt.Println("Zero value for numeric is: 0 (zero)")

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	//%q for a string = a double-quoted string safely escaped with Go syntax

}

func TypeConversion() {
	var i int = 10
	var f float64 = math.Sqrt(float64(i))
	var z uint = uint(f)
	fmt.Println("The expression T(v) converts the value v to the type T, all conversions are explicit")
	fmt.Println(i, f, z)
	fmt.Printf("int %v is converted to a float64, then passed to Sqrt, which returns float64 %v, the returned float64 is converted to uint %v\n", i, f, z)

}
