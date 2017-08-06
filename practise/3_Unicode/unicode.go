package main

import "fmt"
import "github.com/raj-it-geek/go/myutils"

func main() {
	myutils.ClearScreen()

	for teluguUnicode := '\u0C00'; teluguUnicode <= '\u0C7F'; teluguUnicode++ {
		fmt.Printf("\nUnicode: %+5q\tTelugu: %q", teluguUnicode, teluguUnicode)
	}

	fmt.Println()
}
