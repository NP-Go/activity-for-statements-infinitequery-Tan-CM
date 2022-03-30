package main

import "fmt"

func main() {
	//Insert your code here
	// infinite for loop like while (true) {}
	var number int
	var dummy string
	for {
		fmt.Println("Enter a number, -1 to Exit")
		fmt.Scanf("%d", &number)
		fmt.Scanf("%s", &dummy)

		// Exit for any number less than 0
		if number < 0 {
			break
		}

		if number%2 == 1 {
			fmt.Println("Number is Odd")
		} else {
			fmt.Println("Number is Even")
		}
		if number < 10 {
			fmt.Println("Single Digit")
		} else {
			fmt.Println("More than One Digit")
		}
	}
}
