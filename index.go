package main

import (
	"fmt"
)

var fruit = "Mango"
var Rawfruit string = "Raw Mango"

const a = 20

func main() {
	ripeFruit := "Ripe Fruit"
	fmt.Println(Rawfruit)
	Rawfruit = ripeFruit
	fmt.Println(fruit)
	fmt.Println(ripeFruit)
	fmt.Println(Rawfruit)
	fmt.Println("Hello World")
	fmt.Println(20-2)
	fmt.Println(2+2)
	// a = 30
	fmt.Println(a)
}
