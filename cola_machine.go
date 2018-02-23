// Write a program that presents the user w/ a choice of your 5 favorite 
// beverages (Coke, Water, Sprite, ... , Whatever).
// Then allow the user to choose a beverage by entering a number 1-5.
// Output which beverage they chose.

// ★ If you program uses if statements instead of a switch statement, 
// modify it to use a switch statement. If instead your program uses a 
// switch statement, modify it to use if/else-if statements.

// ★★ Modify the program so that if the user enters a choice other than 
// 1-5 then it will output "Error. choice was not valid, here is your money back."

package main 

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Pick a beverage: coke, water, pepsi, fanta, or sprite")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() 

	switch scanner.Text() {
		case "coke": 
			fmt.Printf("You've ordered coke!")
		case "water":
			fmt.Printf("You've ordered water!")
		case "pepsi":
			fmt.Printf("You've ordered pepsi!")
		case "fanta":
			fmt.Printf("You've ordered fanta!")
		case "sprite":
			fmt.Printf("You've ordered sprite!")
	default:
		fmt.Println("Error. choice was not valid, here is your money back.")
	} // end switch
} // end main
