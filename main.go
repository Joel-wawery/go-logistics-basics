package main

import "fmt"

/*creating a struct to model a blueprint of destination as key while the distance in 
km from my warehouse as the value*/
type Destination struct {
	name  string
	distance int
	weather string
    shipments string 
}



func main() {
//code to create the instances of the struct created above by creating a map to the struct
	destinations := map[string]Destination{
		"New York":    {"New York", 300, "Sunny", "Electronics"},
		"Los Angeles": {"Los Angeles", 1200, "Cloudy", "Furniture"},
		"Chicago":     {"Chicago", 800, "Rainy", "Clothing"},
		"Houston":     {"Houston", 1500, "Sunny", "Machinery"},
	}
	
	fmt.Println("Logistics Information: ", destinations)
	//iterating through the map to print each destination and its shipments
	for key, dest := range destinations {
		fmt.Printf("Destination: %s\n", key)
		fmt.Printf("  Shipments: %s\n", dest.shipments)
	}

}