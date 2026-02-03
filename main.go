package main

import "fmt"



func main() {
	//code to create a new map for a fleet of trucks i have for my system
	//the key is the ID of the truck and the value is the current status of the truck
	truckFleet := map[string]string{
		"Truck1": "Available",
		"Truck2": "In Transit",
		"Truck3": "Under Maintenance",
	}	
	//code to add a fleet to the map
	truckFleet["Truck4"] = "Available"
	truckFleet["Truck5"] = "In Transit"
	//code to update the status of the map
	truckFleet["Truck2"] = "has finished delivery of the product available"
	//code to delete a particular truck since its not in operation
	delete(truckFleet, "Truck3")
	//print the final map
	fmt.Println(truckFleet)

	//code to lopp throough the map and print the staus of each truck and inspect through range
	for TruckId , status := range truckFleet {
		fmt.Printf("Truck ID: %s, Status: %s\n",TruckId , status)
	}
}
