package advanced

import (
	"fmt"
	"reflect"
)

func Pointers() {
	fmt.Println("A pointer holds the memory address of a value")
	fmt.Println("The type *T is a pointer to a value of type T")

	x, y := 123, 789
	fmt.Println("x ", x)
	fmt.Println("y ", y)

	var xPointer *int = &x             //& returns the address i.e. pointer.
	var xGetViaPointer int = *xPointer //dereference the *pointer reads the value at the address pointed to.

	fmt.Printf("xPointer address: %p type: %s value: %d\n", xPointer, reflect.TypeOf(xPointer), xGetViaPointer)

	*xPointer = 456 //modify the value at the address pointed to, via the pointer
	fmt.Printf("xPointer address: %p type: %s value: %d\n", xPointer, reflect.TypeOf(xPointer), *xPointer)

	xPointer = &y //pointer reassigned to point at a different value of the same type
	fmt.Printf("xPointer address: %p type: %s value: %d\n", xPointer, reflect.TypeOf(xPointer), *xPointer)

	*xPointer = *xPointer * 10 //manipulate the value at the address pointed to, via the pointer
	fmt.Printf("xPointer address: %p type: %s value: %d\n", xPointer, reflect.TypeOf(xPointer), *xPointer)

}
