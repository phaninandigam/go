
package main

import "fmt"
import "strings"
import "myutils"

func main() {
	var userStatus = "N"

	myutils.ClearScreen()

	fmt.Printf("\n\nWelcome to the Login System!")
	fmt.Printf("\nDo you have a Login Account? (y/n):")
	//fmt.Scanf("%s", &userStatus)
	fmt.Println(userStatus)

	if strings.ToUpper(userStatus) == "N" {
		//fmt.Printf("\nNew User")
		createNewUserAccount()
	}
}

func createNewUserAccount() {
	var userName string

	fmt.Printf("\n\nPlease sign up for a new User Account by entering the details requested below:\n")
	fmt.Printf("\t User Name: ")
	fmt.Scanf("%s", &userName)
	strings.Trim(userName, " ")
	fmt.Printf("\n....%s\n", userName)
}



