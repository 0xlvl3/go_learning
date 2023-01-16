package main

import "log"

func main(){
	var myString string
	myString = "Green"
	
	log.Println(myString)

	// & is a reference to a variable.
	changeUsingPointer(&myString)
	log.Println("After func call myString is set to: ", myString)
}

// * will create a pointer type.
func changeUsingPointer(s *string){
	log.Println("s is set to ", s)
	newValue := "Red"
	*s = newValue
}
