package advanced

import "fmt"
import "reflect"

type Vertex struct {
	X int
	Y int
	z int
}

func Structs() {
	fmt.Println("A struct is a collection of fields.", Vertex{1, 2, 3})
	fmt.Println("Unitialised struct fields have their zero value.", Vertex{})
	fmt.Println("A struct initialised using named fields (any order).", Vertex{z: 11, X: 22, Y: 33})

	v := Vertex{1, 2, 3}
	v.X = 4
	fmt.Println("Struct fields are accessed using '.' operator.", v)

	var vertexPointer = &v
	fmt.Printf("Struct pointer, address: %p type: %s value: %d\n", vertexPointer, reflect.TypeOf(vertexPointer), *vertexPointer)

	fmt.Printf("Struct fields can be accessed via the pointer (vertexPointer.X): %d which is shorthand for dereferencing (*vertexPointer).X: %d\n", vertexPointer.X, (*vertexPointer).X)

	vertexPointer = &Vertex{5, 6, 7}
	fmt.Printf("Struct pointer to the address of a struct literal: %p type: %s value: %d\n", vertexPointer, reflect.TypeOf(vertexPointer), *vertexPointer)

	//next: https://tour.golang.org/moretypes/6
}
