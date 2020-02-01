// maps
package advanced

import (
	"fmt"
)

func Maps() {
	fmt.Println("Maps")

	fmt.Println("nil map")
	var nilMap map[string]Vertex
	fmt.Println("\tAn unititialised map has a zero value of nil (but still has a type of 'map')", nilMap == nil, nilMap, len(nilMap))
	fmt.Println("\titems cannot be added to a nil map")

	fmt.Println("The make function can be used to make initialized and ready to use maps")
	makeMap := make(map[string]Vertex)
	fmt.Println("\tAn ititialised map using the make func", makeMap == nil, makeMap, len(makeMap))
	makeMap["foo"] = Vertex{1, 2, 3}
	fmt.Println("\t", makeMap)
	//var mapStringToVertex map[string]Vertex

	fmt.Println("Map literal")
	var mapLiteral = map[string]Vertex{
		"aaa": Vertex{4, 5, 6},
		"bbb": Vertex{7, 8, 9},
	}
	fmt.Println("\t", mapLiteral)

	fmt.Println("Map literal with type name omitted for elements")
	mapLiteral = map[string]Vertex{
		"aaa":/*Vertex*/ {40, 50, 60},
		"bbb":/*Vertex*/ {70, 80, 90},
	}
	fmt.Println("\t", mapLiteral)

	fmt.Println("Mutating maps")
	mut := make(map[string]int)
	fmt.Println("\tmap", mut)
	mut["a"] = 1
	fmt.Println("\tinserted element", mut)
	fmt.Println("\tget element", mut["a"])
	element, ok := mut["a"]
	fmt.Println("\ttest element", element, ok)
	delete(mut, "a")
	fmt.Println("\tdeleted element", mut)
	element, ok = mut["a"]
	fmt.Println("\ttest element after delete", element, ok)
}
