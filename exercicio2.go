package main

import "fmt"

func Summation(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(Summation(2)) // Saída: 15
	fmt.Println(Summation(8)) // Saída: 55
	fmt.Println(Summation(0)) // Saída: 0
}
