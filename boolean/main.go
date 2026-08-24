package main

import "github.com/01-edu/z01"
import "os"

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr % 2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	arg := os.Args[1:]
	
	if len(arg)% 2 == 0 {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
	
	
}