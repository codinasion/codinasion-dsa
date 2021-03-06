// Write a programme to validate palindrome
package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Print("Enter a string : ")
	fmt.Scan(&str)
	fmt.Println("\nIs '"+str+"' a palindrome : ", isPalindrome(str))
}

func isPalindrome(str string) bool {
	var i, j int
	i = 0
	j = len(str) - 1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
