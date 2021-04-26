package main

import "fmt"

// <Factor> returns a factor except 1 and num.
func Factor(num int) []int {
	result := []int{}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

// <IsPrime> returns if number is a prime number
func IsPrime(num int) bool {
	result := Factor(num)

	if len(result) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	num := 30
	result := []int{}

	for _, i := range Factor(num) {
		if IsPrime(i) {
			result = append(result, i)
		}
	}

	fmt.Println(result)
}
