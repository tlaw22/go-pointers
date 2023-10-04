package main

import "log"

func main() {

	var myString string
	myString = "Green"

	log.Println("My favorite color is: ", myString)
	changePoint(&myString)

	log.Println("My favorite color is: ", myString)

}

func changePoint(s *string) {
	newValue := "Red"
	*s = newValue
	log.Println("The memory location of myString is:", s)
}

// * to passign a pointer
// & to call a pointer
