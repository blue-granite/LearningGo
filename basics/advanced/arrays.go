package advanced

import "fmt"

func Arrays() {
	fmt.Println("The type [n]T is an array of n values of type T")

	var ints [3]int       //zero values
	var strings [2]string //zero values

	fmt.Println("[]int - zero values", ints)
	fmt.Println("[]strings - zero values", strings)

	fmt.Println("first int", ints[0])
	fmt.Println("first string", strings[0])

	ints[0] = 1
	ints[1] = 2
	ints[2] = 3

	strings[0] = "hello"
	strings[1] = "world"

	fmt.Println("[]int", ints)
	fmt.Println("[]strings", strings)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("[]int primes", primes)

	length := len(primes)
	fmt.Println("[]int primes length", length)

	for i, v := range primes {
		fmt.Printf("Range over array, [%d]=%d\n", i, v)
	}
}
