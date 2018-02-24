/* Write a program that asks the user to enter the number of pancakes eaten for breakfast by
10 different people (Person 1, Person 2, ..., Person 10) Once the data has been entered the
program must analyze the data and output which person ate the most pancakes for breakfast.

★ Modify the program so that it also outputs which person ate the
least number of pancakes for breakfast.

★★★★ Modify the program so that it outputs a list in order of the number of pancakes eaten by
all 10 people.

i.e.
Person 4: ate 10 pancakes
Person 3: ate 7 pancakes
Person 8: ate 4 pancakes
...
Person 5: ate 0 pancakes */

package main

import (
	"fmt"
)

func main() {

	var pancakes [10]int
	var number int
	var holder int

	var maxIndex int
	var minIndex int

	// Get user input for the number of pancakes ten people consume
	fmt.Println("Please list the number of pancakes each person eats")

	for i := 0; i < 10; i++ {
		fmt.Scanln(&number)
		pancakes[i] = number
	} // end for

	original := pancakes

	// sort the pancakes by implementing a sorting algorithm
	// in this case it's bubble sort (not efficient for large arrays numbers)
	for z := 0; z < 10; z++ {
		for i := 0; i < 9; i++ {
			j := i + 1
			if pancakes[i] > pancakes[j] {
				holder = pancakes[i]
				pancakes[i] = pancakes[j]
				pancakes[j] = holder
			} // end if
		} // end inner for
	} // end outer for

	// Print the number of pancakes consumed in order
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if original[i] == pancakes[j] {
				fmt.Printf("Person %v ate %v pancakes \n", i+1, pancakes[j])

				switch j {
				case 0:
					minIndex = i
				case 9:
					maxIndex = i
				} // end switch
				j = 10
			} // end if
		} // end inner for
	} // end outer for

	// Print who ate the least and who ate the most pancakes
	fmt.Printf("\n%v is the least number of pancakes consumed by Person %v \n", pancakes[0], minIndex+1)
	fmt.Printf("%v is the most number of pancakes consumed by Person %v \n", pancakes[9], maxIndex+1)
} // end main
