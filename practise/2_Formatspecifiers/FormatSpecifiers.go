package main

import (
	"fmt"

	"github.com/raj-it-geek/go/myutils"
)

func main() {
	var number = 3077

	myutils.ClearScreen()

	for number > 0 {
		// fmt.Printf("\n\nPlease enter a number (0 to exit):")
		// fmt.Scanf("%d", &number)

		if number > 0 {

			// number = 3077

			fmt.Printf("\n     Decimal :  %%d  : %d", number) // %d  : 3077
			fmt.Printf("\n     Binary  :  %%b  : %b", number) // %b  : 110000000101
			fmt.Printf("\n     Octal   :  %%o  : %o", number) // %o  : 6005

			fmt.Println()
			fmt.Printf("\n Hexadecimal :  %%x  : %x", number)  // %x  : c05
			fmt.Printf("\n Hexadecimal :  %%#x : %#x", number) // %X  : 0Xc05
			fmt.Printf("\n Hexadecimal :  %%X  : %X", number)  // %x  : C05
			fmt.Printf("\n Hexadecimal :  %%#X : %#X", number) // %X  : 0XC05

			fmt.Println()
			fmt.Printf("\n\t UTF :  %%U  : %U", number)  // %U  : U+0C05
			fmt.Printf("\n\t UTF :  %%+q : %+q", number) // %+q : '\u0c05'
			fmt.Printf("\n\t UTF :  %%q  : %q", number)  // %q  : 'అ'
			fmt.Printf("\n\t UTF :  %%c  : %c", number)  // %c  : అ

			fmt.Println("\n\n===========================================")
			fmt.Printf("\n Generic Placeholder %%v : %v", number) // %v : 3077
			fmt.Printf("\n %%#v : %#v", number)                   // 3077

			// Width Specifiers
			fmt.Println()
			fmt.Printf("\nDefault			:  %%d   : %d", number)         // %d  : 3077
			fmt.Printf("\nRight Justified 	:  %%7d  : %7d", number) // %d  :    3077
			fmt.Printf("\nRight Justified 	:  %%7d  : %7d", 16)     // %d  :      16
			fmt.Printf("\nLeft Justified  	:  %%-d  : %-d", number) // %d  :16

			fmt.Println()
			fmt.Printf("\nDefault				:  %%0d  : %0d", number)                  // %d  : 3077
			fmt.Printf("\nZero-Fill Right Justified 	:  %%07d : %07d", number) // %d  :    3077
			fmt.Printf("\nZero-Fill Right Justified 	:  %%07d : %07d", 16)     // %d  :      16
			fmt.Printf("\nZero-Fill Left Justified  	:  %%-0d : %-0d", number) // %d  :16

			number = 0
		}
	}
	fmt.Printf("\n\n")
}
