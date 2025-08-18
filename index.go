package main

import (
	"fmt"
	"strings"
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
	fmt.Println(20 - 2)
	fmt.Println(2 + 2)
	// a = 30
	fmt.Println(a)
	fmt.Printf("%c", ripeFruit[1])
	fmt.Println()
	msg := "one"
	msg2 := "One"
	fmt.Println(strings.Compare(msg, msg2))
	// arrays
	var arr1 = [6]int{1, 2, 3, 4, 5}
	arr1[0] = 69
	fmt.Println(arr1)
	// slices
	var arr2 = []int{1, 2, 3, 4, 5}
	fmt.Println(arr2[:4])
	fmt.Println(arr2[2:4])
	fmt.Println(arr2[3:])

	

}
