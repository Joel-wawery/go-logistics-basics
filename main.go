package main

import "fmt"


//code to create a struct for our trucking system(blueprint for our system)
	type Truck struct {
		ID  string
		status  string
		capacity  int
		currentLoad  float64
		location  string

	}


func main() {
	//creating a real truck and filling up its details 
	truck1 := Truck{
		ID : "T001",
		status : "available",
		capacity : 10000,
		currentLoad : 0.0,
		location : "Warehouse A",
	}
	fmt.Println(truck1)

	fleet := make(map[Truck]Truck)
	fleet[truck1] = Truck {
		ID : "T002",
		status : "in transit",
		capacity : 15000,
		currentLoad : 5000.0,
		location : "On Route to City B",
	}
   	fmt.Println(fleet)

	//looping through the fleet of trucks and printing their details
	for key, value := range fleet {
		fmt.Println("Truck Key:", key)
		fmt.Println("Truck Value:", value)
	}
	
}