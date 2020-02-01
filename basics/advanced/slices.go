// slices
package advanced

import (
	"fmt"
	"reflect"
	"strings"
)

//https://blog.golang.org/go-slices-usage-and-internals

func Slices() {
	fmt.Println("A slice is a dynamically-sized view of an array.")
	fmt.Println("The type []T is a slice with elements of type T.")
	fmt.Println("A slice is formed by specifying two indices, a low (included) and high (excluded) bound.")

	array := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("array", array)

	var slice []int = array[1:4]
	fmt.Println("slice, array[1:4]", slice)

	fmt.Println("When slicing, you may omit the high or low bounds to use their defaults instead.")
	fmt.Println("The default is zero for the low bound and the length of the slice for the high bound.")
	var sliceNoLowBound []int = array[:4]
	var sliceNoHighBound []int = array[2:]
	var sliceNoBounds []int = array[:]
	fmt.Println("sliceNoLowBound, array[:4]", sliceNoLowBound)
	fmt.Println("sliceNoHighBound, array[2:]", sliceNoHighBound)
	fmt.Println("sliceNoBounds, array[:]", sliceNoBounds)

	fmt.Println("A slice does not store any data, it just references a section of an underlying array.")
	fmt.Println("Changing the elements of a slice modifies the corresponding elements of its underlying array.")

	var anotherSlice []int = array[0:3]

	slice[0] = 99
	fmt.Println("slice, element 0 of slice modified to 99", slice)
	fmt.Println("anotherSlice, other slices that share the array will see changes", anotherSlice)
	fmt.Println("array, element 0 of slice corresponds to element 1 of array ", array)

	fmt.Println("A slice literal is like an array literal without the length.")
	sliceLiertal := [] /*no length*/ bool{true, false, true, true, false, true}
	fmt.Println("Type of a slice doesn not include length", sliceLiertal, reflect.TypeOf(sliceLiertal))

	fmt.Println("A slice has both a length and a capacity.")
	fmt.Println("The length of a slice is the number of elements it contains.")
	fmt.Println("The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.")

	fmt.Println("slice, length", len(slice), slice)
	fmt.Println("slice, capacity", cap(slice), slice)

	fmt.Println("The length and capacity of a slice can be modified by resliced - assuming sufficient capacity.")
	slice = slice[2:5]
	fmt.Println("slice, length", len(slice), slice)
	fmt.Println("slice, capacity", cap(slice), slice)

	var nilSlice []int
	fmt.Println("An unititialised slice has a zero value of nil (but still has a type of 'slice')", nilSlice == nil, nilSlice, len(nilSlice), cap(nilSlice))

}

func SlicesOfSlices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	fmt.Println("Board", board)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"

	fmt.Println("Board elements", board)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func SlicesMake() {
	fmt.Println("Slices can be created with the built-in make function.")
	sliceLen5Cap5 := make([]int, 5) // len=5 cap=5
	printSlice("sliceLen5Cap5", sliceLen5Cap5)

	sliceLen5Cap9 := make([]int, 5, 9) // len=5 cap=9
	printSlice("sliceLen5Cap9", sliceLen5Cap9)

}

func SlicesAppend() {
	fmt.Println("It is common to append new elements to a slice, and so Go provides a built-in append function.")

	var slice []int
	printSlice("An 'nil' slice", slice)

	fmt.Println("Append works on 'nil' slices.")
	for i := 0; i < 5; i++ {
		slice = append(slice, i)
		printSlice("\tappend 'i' to slice", slice)
	}

	fmt.Println("Append more than one element.")
	slice = append(slice, 5, 6, 7, 8, 9)
	printSlice("\tappend multiple elements to slice", slice)

	fmt.Println("If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.")
	var array = [5]int{2, 3, 5}
	fmt.Println("\tarray", array)
	var arrayBackedSlice = array[:3]
	printSlice("\tarrayBackedSlice", arrayBackedSlice)

	arrayBackedSlice = append(arrayBackedSlice, 7, 11)
	fmt.Println("\tarray", array)
	printSlice("\tarrayBackedSlice + 7, 11", arrayBackedSlice)

	arrayBackedSlice = append(arrayBackedSlice, 13)
	fmt.Println("\tarray, doesn't contain 13", array)
	printSlice("\tarrayBackedSlice + 13 as capacity was exceeded a new backing array used by the slice", arrayBackedSlice)

	fmt.Println("The slice returned from append op that exceeds the capacity a new slice and backing array are created")
	var startingSlice = make([]int, 2, 4)
	printSlice("\tstartingSlice", startingSlice)
	var appendedSliceUnderCapacity = append(startingSlice, 99)
	printSlice("\tappendedSliceUnderCapacity", appendedSliceUnderCapacity)
	var appendedSliceOverCapacity = append(appendedSliceUnderCapacity, 1, 2, 3)
	printSlice("\tappendedSliceOverCapacity", appendedSliceOverCapacity)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v    (pntr %v)\n",
		s, len(x), cap(x), x, reflect.ValueOf(x).Pointer())
}
