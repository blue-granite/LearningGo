package flow

import (
	"fmt"
)

func For() {
	//The init statement: executed before the first iteration
	//The condition expression: evaluated before every iteration
	//The post statement: executed at the end of every iteration
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func ForConidition() {
	//The init and post statements are optional.
	// Making it a while loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func ForEver() {
	//Forever loop
	for {
		// fmt.Println("forever")
	}
}
