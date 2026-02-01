/*in Go every file has to have a package declartion so that it shows each file is in a package
main is special as it tells the Go language that hey this is an executable program not a library
main thing to note is in Go every single file belongs to a package.....
in hindsight this program belongs to the main package and uses the fmt package starts running
in the main func and prints out hello world in the terminal
*/
package main

/* now we import the format package which we will use to print out fucntions ,format texts and
reading inputs ...so without this the println ,text formatting and input reading wont happen
*/
import "fmt"

/* why write the main function ? its because the main is the entry point of the program
if we dont have the entry point of the program then Go wont run the program at all because ther
is nowhere to go into ..so basically its the main entry point of any program to be written */
//go has semicolon insertion rule which basically adds a semicolon after the end of every brace

func main() {
	//variable to show the capacity of each truck 
	var truckCapacity int = 5000 // in kgs
	//variable to show the current load available
	var currentLoad float64 = 1200.50 // in kgs 
	/* we are going to add both to get the total capacity and since Go does not allow types to mix ,
	we are going to convert truckcapacity to float64 so that they are of the same tyep to be added to get the total */

	totalWeight := float64(truckCapacity) + currentLoad

	//print out the totalWeight of the truck
	fmt.Println("the total weight of the truck is ", totalWeight, "kgs")

}
/* ultimately the main reason why we are using variables here is to first create a
 box in the memory so that we can store value in it and then afterwards we choose meaningful names 
 to the box so that we know the function of the box and then we specify the type of the 
 box we have created a.k.a the variable so that we know what type of data we are going to store and use
  in that box that we have given meaningful name to  */

/* it is also bad practice to declare variables outside of the main function because its
accessible from anywhere and also it can be modified from anywhere hence making it dangourous
dur to silent bugs created  */  