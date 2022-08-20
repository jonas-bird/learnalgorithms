package main

import (
	"fmt"
)

func main() {
    var numbers []int = []int{5, 4, 2, 3, 1, 0}
    fmt.Println("Our list of numbers is:", numbers)

    sweep(numbers)

}

func sweep(numbers []int) {
	// set N to the size of the list
    var N int = len(numbers)

	// initialize variables for the indexs of the values we are
	// comparing and assign them starting values of 0 and 1
    var firstIndex int = 0
    var secondIndex int = 1

	// Loop over the index
    for secondIndex < N {

		// printf testing code
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		fmt.Println("Comparing the following: ", firstNumber, secondNumber)

		firstIndex++
		secondIndex++
	}
}
