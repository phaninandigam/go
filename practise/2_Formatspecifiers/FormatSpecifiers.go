package main

import "fmt"
import "github.com/raj-it-geek/go/myutils"

func main() {
	integerNumber := 3077
	floatNumber := 123456789.123456789
	booleanFlag := true

	myutils.ClearScreen()

	for integerNumber > 0 {
		// fmt.Printf("\n\nPlease enter a Number (0 to exit):")
		// fmt.Scanf("%d", &integerNumber)

		if integerNumber > 0 {

			// integerNumber = 3077

			fmt.Printf("\n     Decimal :  %%d  : %d", integerNumber) // %d  : 3077
			fmt.Printf("\n     Binary  :  %%b  : %b", integerNumber) // %b  : 110000000101
			fmt.Printf("\n     Octal   :  %%o  : %o", integerNumber) // %o  : 6005

			fmt.Println()
			fmt.Printf("\n Hexadecimal :  %%x  : %x", integerNumber)  // %x  : c05
			fmt.Printf("\n Hexadecimal :  %%#x : %#x", integerNumber) // %X  : 0Xc05
			fmt.Printf("\n Hexadecimal :  %%X  : %X", integerNumber)  // %x  : C05
			fmt.Printf("\n Hexadecimal :  %%#X : %#X", integerNumber) // %X  : 0XC05

			fmt.Println()
			fmt.Printf("\n\t UTF :  %%U  : %U", integerNumber)  // %U  : U+0C05
			fmt.Printf("\n\t UTF :  %%+q : %+q", integerNumber) // %+q : '\u0c05'
			fmt.Printf("\n\t UTF :  %%q  : %q", integerNumber)  // %q  : 'అ'
			fmt.Printf("\n\t UTF :  %%c  : %c", integerNumber)  // %c  : అ

			// str := '\u0950'
			// fmt.Printf("\n\t UTF :  %%q  : %q", str) // %q  : 'ॐ'
			// fmt.Printf("\n\t UTF :  %%c  : %c", str) // %c  : ॐ

			fmt.Println()
			fmt.Printf("\n Scientific Notation :  %%f  : %f", floatNumber) // %f  : 123456789.123457
			fmt.Printf("\n Scientific Notation :  %%F  : %F", floatNumber) // %F  : 123456789.123457
			fmt.Printf("\n Scientific Notation :  %%e  : %e", floatNumber) // %e  : 1.234568e+08
			fmt.Printf("\n Scientific Notation :  %%E  : %E", floatNumber) // %E  : 1.234568E+08
			fmt.Printf("\n Scientific Notation :  %%g  : %g", floatNumber) // %g  : 1.2345678912345679e+08
			fmt.Printf("\n Scientific Notation :  %%G  : %G", floatNumber) // %G  : 1.2345678912345679e+08

			// Width Specifiers
			fmt.Println("\n\n===========================================")
			fmt.Printf("\n Default                            :  %%d   : %d", integerNumber)  // %d  : 3077
			fmt.Printf("\n At least 7-wide, Right Justified   :  %%7d  : %7d", integerNumber) // %7d  :    3077
			fmt.Printf("\n At least 7-wide, Right Justified   :  %%7d  : %7d", 16)            // %7d  :      16
			fmt.Printf("\n At least 7-wide, Left Justified    :  %%-d  : %-d", integerNumber) // %-d  :16

			fmt.Println()
			fmt.Printf("\n default Width, default Precision:  %%f      : %f", floatNumber)      // %f  : 123456789.123457
			fmt.Printf("\n Width 9, default Precision 	 :  %%9f     : %9f", floatNumber)        // %9f  : 123456789.123457
			fmt.Printf("\n default Width, Precision 2 	 :  %%.2f    : %.2f", floatNumber)       // %.2f  : 123456789.12
			fmt.Printf("\n Width 9, Precision 2      	 :  %%9.2f   : %9.2f", floatNumber)       // %9.2f  : 123456789.12
			fmt.Printf("\n Width 9, Precision 0  	         :  %%9.f    : %9.f", floatNumber)    // %9.f  : 123456789
			fmt.Printf("\n Width 14, Precision 2      	 :  %%14.2f  : %14.2f", floatNumber)     // %14.2f  :   123456789.12
			fmt.Printf("\n Zero-Padded, Width 14, Precsn 2 :  %%014.2f : %014.2f", floatNumber) // %014.2f  : 00123456789.12

			fmt.Println("\n\n===========================================")
			fmt.Printf("\n Generic Placeholder %%v : %v", integerNumber) // %v : 3077
			fmt.Printf("\n %%#v : %#v", integerNumber)                   // %#v: 3077

			fmt.Println("\n\n===========================================")
			fmt.Printf("\n Value type : %%T : %T", integerNumber)     // int
			fmt.Printf("\n Value type : %%T : %T", floatNumber)       // float64
			fmt.Printf("\n Value type : %%T : %T", booleanFlag)       // bool
			fmt.Printf("\n\n Boolean Word :  %%t  : %t", booleanFlag) // %t : true

			integerNumber = 0
		}
	}
	fmt.Printf("\n\n")
}
