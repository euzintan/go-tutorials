package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var age uint
	var livesLeft uint = 3

	CheckIfGameOver := func() bool {
		if livesLeft <= 0 {
			fmt.Println("Womp Womp Game Over")
		} else {
			fmt.Printf("You just lost 1 life, you have %v lives left\n", livesLeft)
		}
		return (livesLeft <= 0)
	}

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Welcome to my game, %v\nEnter your age: ", name)
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Printf("%v?! You're old, but wise\n", age)
	} else {
		fmt.Printf("%v?! too young!\nCome back in %v years\n", age, 10-age)
		return
	}

	var answer string
	var ended bool = false

	for !ended {
		fmt.Println("What do you hear first, Lightning or Thunder?")

		fmt.Scan(&answer)

		if strings.ToUpper(answer) == "LIGHTNING" {
			fmt.Println("Wrong, you can't hear lightning")
			livesLeft--
			ended = CheckIfGameOver()
		} else if strings.ToUpper(answer) == "THUNDER" {
			fmt.Println("Correct!\nCongratulations I'm out of questions for you!")
			return
		} else {
			fmt.Println("That's not one of the options")
			livesLeft--
			ended = CheckIfGameOver()
		}
	}

}
