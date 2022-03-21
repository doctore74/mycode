/* Wrt - 2 functions to print my name and address */
   
package main

import (
    "fmt"
)

var city, state = "Duesseldorf", "NRW"


func printName(firstName, lastName string) string {
	fullName := "My name is: " + firstName + " " + lastName
	return fullName
}


func main() {
    fmt.Println(printName("Christian", "Wirtz"))
    fmt.Println(city, state)
}

