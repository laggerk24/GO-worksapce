package main

import (
	"fmt"
	"strings"
)

var fruit = "Mango"
var Rawfruit string = "Raw Mango"

const a = 20

type Person struct {
	name    string
	salary  int
	address string
}

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
	fmt.Println("Arrays")
	for num := range arr1 {
		fmt.Print(num)
		fmt.Print(",")
	}
	// slices
	var arr2 = []int{8, 9, 10, 11, 12}
	fmt.Println(arr2[:4])
	fmt.Println(arr2[2:4])
	fmt.Println(arr2[3:])

	var userData = map[string]int{
		"lagger": 24,
		"rahul":  22,
	}
	userData["ram"] = 23
	fmt.Println(userData["sam"])

	for key, value := range userData {
		fmt.Println(key + " ")
		fmt.Print(value)
		fmt.Println()
	}

	var object Person
	object.name = "Lagger"
	object.salary = 9000000
	object.address = "Hyderabad"
	fmt.Println("Hello my name is", object.name, "living in", object.address, "with salary of", object.salary)
}
