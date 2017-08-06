package main

import "fmt"
import "github.com/raj-it-geek/go/myutils"

func main() {
	var number = 1

	myutils.ClearScreen()

	for number > 0 {
		fmt.Printf("\nPlease enter a number (0 to exit):")
		fmt.Scanf("%d", &number)
		if number > 0 {
			fmt.Printf("\n Decimal: %d\n Binary\t: %b\n Octal\t: %o\n Hexadecimal: %#x\n\n", number, number, number, number)
		}
	}
	fmt.Println()
}
