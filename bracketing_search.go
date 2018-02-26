/* Write a program that calculates a random number 1 through 100. The program then asks the user 
to guess the number. If the user guesses too high or too low then the program should output 
"too high" or "too low" accordingly. The program must let the user continue to guess until the user 
correctly guesses the number.

★ Modify the program to output how many guesses it took the user to correctly guess the right number.

★★ Modify the program so that instead of the user guessing a number the computer came up with, 
** the computer guesses the number that the user has secretely decided. The user must tell the computer 
** whether it guesed too high or too low.

★★★★ Modify the program so that no matter what number the user thinks of (1-100) 
**** the computer can guess it in 7 or less guesses. */

package main

import (
	"fmt"
)

func main () {

	var number int
	guess := 50

	fmt.Println("Please input a number between 1-100")
	fmt.Scanln(&number)

	for i := 0; i <= 7; i++ {
		fmt.Println(guess)
		if guess == number {
			fmt.Printf("Your number was %v", guess)
			fmt.Printf("This took %v guesses", i+1)
			i = 7
		} else {
			if number < guess {
				// placeholder

			} else {
				// placeholder
			} // end inner else 

		} // end outer else
	} // end for
} // end main
