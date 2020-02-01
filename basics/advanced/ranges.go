// ranges
package advanced

import (
	"fmt"
)

func Ranges() {

	array := [6]int{2, 3, 5, 7, 11, 13}
	slice := array[1:5]

	fmt.Println("Range over array")
	for i, v := range array {
		fmt.Printf("\t[%d]=%d\n", i, v)
	}
	fmt.Println("Range over slice")
	for i, v := range slice {
		fmt.Printf("\t[%d]=%d\n", i, v)
	}

	var ssmap = map[string]string{
		"a": "AAA",
		"b": "BBB",
	}

	fmt.Println("Range over slice")
	for key, val := range ssmap {
		fmt.Printf("\t[%s]=%s\n", key, val)
	}

	fmt.Println("Range, using only index")
	for i, _ := range array {
		fmt.Printf("\t[%d]\n", i)
	}

	fmt.Println("Range, using only index, omitting _")
	for i := range array {
		fmt.Printf("\t[%d]\n", i)
	}

	fmt.Println("Range, using only value")
	for _, v := range array {
		fmt.Printf("\tv=%d\n", v)
	}

	fmt.Println("Range, using only index, omitting _")
	for p, c := range "≪string≫" {
		fmt.Printf("\tcharacter %#U starts at byte position %d\n", c, p)
	}
}
