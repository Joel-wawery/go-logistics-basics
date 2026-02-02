package main

import "fmt"



func main() {
	//code to create a slice where we are expecting 5 orders but it can grow or shrink
	orders := make([]float64, 0 , 5)

	//adding data to the slice and here the append checks capacity if its full it will allocate a bigger memory

	orders = append(orders, 120.5 ,450.5 , 789.5)

	fmt.Println("Orders: ", orders)
	fmt.Println("Length: ", len(orders))
	fmt.Println("Capacity: ", cap(orders))

	firstTwoOrders := orders[:2]
	fmt.Println("First two orders: ", firstTwoOrders)

}
