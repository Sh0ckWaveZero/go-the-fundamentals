package main

import "fmt"

func main() {
	s := []string{"Hello", "world"}
	variadicString(s...)
}

func variadicString(str ...string) {
	for _, s := range str {
		fmt.Println(s)
	}
}
