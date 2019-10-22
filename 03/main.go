package main

import "fmt"

type handle func(int, int) int

func Max(a, b int) int {

	return a + b
}
func haha(a, b int, f handle) int {
	return f(a, b)
}

func main() {

	fmt.Println(haha(2, 3, Max))
}
