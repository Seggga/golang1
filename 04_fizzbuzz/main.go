package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		fizzBuzz := true
		if i%3 == 0 {
			fmt.Print("Fizz")
			fizzBuzz = false
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
			fizzBuzz = false
		}

		if fizzBuzz {
			fmt.Print(i)
		}

		fmt.Println()

	}
	fmt.Println("hello")
}
