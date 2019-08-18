package main

import "fmt"

func main() {
	var coinflip = "heads"
	switch coinflip {
	case "heads":
		fmt.Println("heads")
	case "tails":
		fmt.Println("tails")
	default:
		fmt.Println("landed on edge!")
	}
}
