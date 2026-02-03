package main

import "fmt"



func main() {
	//code to prepare the list of the number of books we have
	booksAvailable := make([]int ,0 , 9)
	
	//code to append the delivery times that have already been made
	booksAvailable = append(booksAvailable, 5 , 13 , 16 , 17 , 19 , 145 ,13545)
	//code to print out the deliveries made so far
	fmt.Println("All deliveries made:", booksAvailable)
	
	//code to slice the list to show the first 2 deliveries made
	booksAvailed := booksAvailable[:5]
	
	//code to print out the first two deliveries made
	fmt.Println("all the main books availed:", booksAvailed)
}
