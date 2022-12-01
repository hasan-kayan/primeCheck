package main

import "fmt"

func main() {
	for {
		fmt.Println("Enter a number to check if it is prime or not")
		var number int
		fmt.Scanln(&number)
		if number < 2 {
			fmt.Println("There is not any prime number lower than 2. ")
		} else {
			var isPrime bool = true
			for i := 2; i < number; i++ {
				if number%i == 0 {
					isPrime = false
					break
				}
			}
			if isPrime == true {
				fmt.Println(number, "is a prime number")
			} else {
				fmt.Println(number, "is not a prime number")
			}
		}
	}
}
