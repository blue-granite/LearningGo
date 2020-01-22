package main

import (
	"fmt"
)

func NoMixPkgInDir() {
	fmt.Println("You can't mix package declarations in the same folder, i.e. this func is in the 'main' package becuase there is already another go file in the directory")
}
