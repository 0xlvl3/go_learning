// Main must be declared first.
package main

// fmt == format.
import "fmt"

// This is called a package level variable
var myName string 

// Must have main function. 
func main() {
	fmt.Println("Hello, world!")
	
	// Variables are strict
	var whatToSay string = "Saying this"
	var i int64 = 4

	fmt.Println(whatToSay)
	fmt.Println(`This is a string?`, i)

	// := is shorthand to tell our variable to take the value from returned from the function
	// and store it within the variable.
	stringReturn1, stringReturn2 := saySomething()
	fmt.Println(`The function returned`, stringReturn1, stringReturn2)
}

// go can return two items if they are declared
func saySomething() (string, string) {

	return `something something`, `!! it also returned another string`
}
