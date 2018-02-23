// Write a program that continues to asks the user to enter any number other than 5 until 
// the user enters the number 5. Then tell the user "Hey! you weren't supposed to enter 5!" 
// and exit the program.

// ★ Modify the program so that after 10 iterations if the user still hasn't entered 5 will 
// tell the user "Wow, you're more patient then I am, you win." and exit.

// ★★ Modify the program so that it asks the user to enter any number other than the number 
// equal to the number of times they've been asked to enter a number. (i.e on the first 
// iteration "Please enter any number other than 0" and on the second iteration 
// "Please enter any number other than 1"m etc. etc. The program must behave accordingly 
// exiting when the user enters the number they were asked not to.)

package main 

import (
	"fmt"
)

func main () {
	var input int

	for i := 0; i < 10; i++ {
		fmt.Printf("Please input any number except for: %v\n", i)
		fmt.Scanln(&input)

		// If the user inputs the index before ten turns, exit the program
		if input == i && i != 10 {
			fmt.Printf("Hey, you weren't suppposed to enter %v!\n", i)
			i = 10
		} else if input != i && i == 9 {
			fmt.Println("Wow, you're more patient than I am, you win.")
		} // end else if
	} // end for
} // end main
