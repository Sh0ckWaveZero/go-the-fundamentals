package main

import "fmt"

func main() {
	// fmt.Println("Hello, world.")

	// name := "Bob"
	// appendGreeting(&name)
	// fmt.Println(name)
	abc()
	efg()
	cde()
}

func greeting(name string) string {
	return "hello, " + name
}

func isOdd(n int) bool {
	if n%2 != 0 {
		return true
	} else {
		return false
	}
}

func sumOfFirst(n int) int {
	result := 0
	for i := 0; i <= n; i++ {
		result += i
	}
	return result
}

// A prime number is a number greater than 1 with only two factors – themselves and 1
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func appendGreeting(s *string) {
	// *s = "Hi," + *s
	*s = fmt.Sprintf("Hi, %s", *s)
}

func abc() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "apple", "banana" and "coconut" needed

	fmt.Print(s[:3])
	// [apple banana coconut]
}

func efg() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "elderberries", "figs" and "grapes" needed

	fmt.Print(s[4:])
	// * [elderberries figs grapes]
}

func cde() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "coconut", "durian" and "elderberries" needed

	fmt.Print(s[2:5])
	// * [coconut durian elderberries]
}
