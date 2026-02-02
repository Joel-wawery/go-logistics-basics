package main

import "fmt"

const maxShillingsGoal int = 20000000

func main() {
	//variable to store the amount of shillings each job pays
	var eachPayingJob int = 1000
	//variable to store the number of jobs needed to reach the goal
	numberOfJobs := maxShillingsGoal / eachPayingJob
	fmt.Println(numberOfJobs)

}
