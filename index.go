package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var fruit = "Mango"
var Rawfruit string = "Raw Mango"

const a = 20

type Person struct {
	name    string
	salary  int
	address string
}

type Employee struct {
	person        Person
	id            string
	officeAddress string
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

	// Struct

	var object Person
	object.name = "Lagger"
	object.salary = 9000000
	object.address = "Hyderabad"
	fmt.Println("Hello my name is", object.name, "living in", object.address, "with salary of", object.salary)

	emp := Employee{
		person: Person{
			name:    "Rahul",
			salary:  900000000,
			address: "Hyderabad",
		},
		id:            "180222",
		officeAddress: "Hyderabad",
	}

	fmt.Println(emp)

	// Inline Struct

	student := struct {
		firstName  string
		secondName string
	}{
		firstName:  "lagger",
		secondName: "king",
	}

	fmt.Println(student)
	test("lagger", 11, 2, 0, 4, 5, 0, 6, 10, 7, 9)
	fmt.Println()
	getAllNames("lagger", "Rahul", "aman")

	fmt.Println(addFunction(4, 5))

	fmt.Println("Anonymous function ")
	testFunc := func() {
		fmt.Println("Expression")
	}
	testFunc()
	add := func(num1 int, num2 int) int {
		return num1 + num2
	}

	fmt.Println("Anonymous function output", add(3, 2))

	fmt.Println("Function with a struct")
	fmt.Println(object.allDetails())

	fmt.Println("Callback FUnction")
	fmt.Println(addViaCallBack(69, 1, add))
	fmt.Println(addViaCallBack(69, 2, addFunction))
	fmt.Println(addViaCallBack(69, 2, func(num1 int, num2 int) int {
		return num1 + num2
	}))

	// Defer Keyword
	fmt.Println("Defer Keyword")
	greet()
	greetWithDefer()

	// Pointers
	x := 35
	var pointerToX *int = &x
	fmt.Println(pointerToX)
	fmt.Println("address of pointer", &pointerToX)
	fmt.Println("Value at pointer", *pointerToX)

	// panic --> like throw of other programs
	ageVerification := func(age int) {
		if age > 60 {
			panic("Retirement age")
		}
	}
	ageVerification(24)
	ageVerification(60)

	// interface in Go
	fmt.Println(log("test 123", DebugLogger{loggerType: "Debug"}))
	fmt.Println(log("test 124", ErrorLogger{loggerType: "Error"}))

	// go Routine
	go goRoutine()
	// time.Sleep(1 * time.Second)

	// channels
	fmt.Println()
	ch := make(chan int) // --> this is unnbuffered by default
	go sendData(ch)
	val := <-ch
	fmt.Println("Channle data recieved", val)

	// Unbuffered channel --> capacity of 0, syncronous communication b/n sender and reciever
	// buffered channel --> specific capacity > 0
	unBufferedCh := make(chan int)
	bufferedCh := make(chan string, 2)
	go func() {
		bufferedCh <- "Ram Ram "
		bufferedCh <- "Saryana"
		close(bufferedCh)
	}()
	go func() {
		unBufferedCh <- 23
	}()

	value := <-unBufferedCh
	fmt.Println("Recieved from unbuffered channel", value)
	valueBuffered := <-bufferedCh
	fmt.Println("Recieved from buffered channel", valueBuffered)
	// valueBuffered = <-bufferedCh
	// fmt.Println("Recieved from buffered channel", valueBuffered)
	for msg := range bufferedCh {
		fmt.Println("Yo")
		fmt.Println("Recieved from buffered channel", msg)
	}

	// WAIT group
	var wg sync.WaitGroup
	wg.Add(3)
	go printMessage(&wg, "Lagger", 2)
	go printMessage(&wg, "Rahul", 5)
	go printMessage(&wg, "Aman", 6)
	wg.Wait()
	fmt.Println("waiting complete ")

	// SELECT
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(6 * time.Second)
		ch1 <- "Message for ch1"
	}()

	go func() {
		time.Sleep(7 * time.Second)
		ch2 <- "Message for ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Recieved", msg1)
	case msg2 := <-ch2:
		fmt.Println("Recieved", msg2)
	case <-time.After(5 * time.Second):
		fmt.Println("nothing received")
	}

}

func (p Person) allDetails() string {
	return p.name + " " + strconv.Itoa(p.salary) + " " + p.address
}

func test(name string, numbers ...int) {
	fmt.Println(name, numbers)
	for num := range numbers {
		fmt.Print(numbers[num], ",")
	}
}

func getAllNames(names ...string) {
	fmt.Println(names)
	for name := range names {
		fmt.Print(names[name], ",")
	}
}

func addFunction(num1 int, num2 int) int {
	return num1 + num2
}

func addViaCallBack(num1 int, num2 int, callback func(int, int) int) int {
	return callback(num1, num2)
}

func greet() {
	fmt.Println("World")
	fmt.Println("Hello")
}

func greetWithDefer() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

type Logger interface {
	logMessage() string
	// sayHi() string
}

type ErrorLogger struct {
	loggerType string
}

type DebugLogger struct {
	loggerType string
}

func (e ErrorLogger) logMessage() string {
	return e.loggerType
}

func (d DebugLogger) logMessage() string {
	return d.loggerType
}

func log(message string, logger Logger) string {
	return logger.logMessage() + " " + message
}

func goRoutine() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func sendData(ch chan<- int) {
	time.Sleep(100 * time.Millisecond)
	ch <- 29
}

func printMessage(wg *sync.WaitGroup, message string, seconds int) {
	defer wg.Done()
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println(message)
}
