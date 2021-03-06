// Grading Program
// Write a program that allows a user to enter the grade scored in a programming class (0-100).
// If the user scored a 100 then notify the user that they got a perfect score.
// * Modify the program so that if the user scored a 90-100 it informs the user they scored an A
// ** Modify the program so that it will notify the user of their letter grade
// ** 0-59 F 60-69 D 70-79 C 80-89 B 90-100 A

package main

import(
	"fmt"
)

func main() {
	var grade int

	fmt.Printf("Please enter your grade: ")
	fmt.Scanln(&grade)

	if grade <= 59 {
		fmt.Println("You received an F")
	} else if grade >= 60 && grade <= 69 {
		fmt.Println("You recieved a D")
	} else if grade >= 70 && grade <= 79 {
		fmt.Println("You received a C") 
	} else if grade >= 80 && grade <= 89 {
		fmt.Println("You received a B")
	} else {
		fmt.Println("You received an A!")
		if grade == 100 {
			fmt.Println("You received a perfect score!")
		} // end if
	} // end else
}
